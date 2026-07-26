"""出荷データの保管(デモなのでインメモリ)。"""


class ShipmentStore:
    def __init__(self) -> None:
        self.shipments: dict[str, dict] = {}

    def create(self, order_id: str) -> dict:
        shipment = {"id": f"shp_{order_id}", "order_id": order_id, "status": "created"}
        self.shipments[shipment["id"]] = shipment
        return shipment

    def get(self, shipment_id: str) -> dict:
        return self.shipments.get(shipment_id, {})

    def update_status(self, shipment_id: str, status: str) -> None:
        if shipment_id in self.shipments:
            self.shipments[shipment_id]["status"] = status
