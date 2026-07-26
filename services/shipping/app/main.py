"""shipping サービス(FastAPI)。出荷の起票・追跡と、運送会社 webhook の受信。"""

from fastapi import APIRouter, FastAPI

from app.carrier import CarrierClient
from app.store import ShipmentStore

app = FastAPI()
router = APIRouter(prefix="/shipments")
store = ShipmentStore()
carrier = CarrierClient()


@router.post("")
async def create_shipment(order_id: str) -> dict:
    """注文から出荷を起票し、運送会社に集荷を依頼する。"""
    shipment = store.create(order_id)
    await carrier.request_pickup(shipment["id"])
    return shipment


@router.get("/{shipment_id}")
async def get_shipment(shipment_id: str) -> dict:
    """出荷状況を返す。"""
    return store.get(shipment_id)


@app.post("/webhooks/carrier")
async def carrier_webhook(payload: dict) -> dict:
    """運送会社からの配送ステータス通知(外部 → 自システム)。"""
    store.update_status(payload.get("shipment_id", ""), payload.get("status", ""))
    return {"ok": True}


app.include_router(router)
