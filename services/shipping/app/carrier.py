"""運送会社 API(外部 SaaS)への HTTP クライアント。"""

import os

import httpx


class CarrierClient:
    def __init__(self) -> None:
        self.token = os.getenv("CARRIER_API_TOKEN", "")

    async def request_pickup(self, shipment_id: str) -> None:
        """集荷を依頼する(自システム → 外部)。"""
        httpx.post("https://api.carrier.example.com/v2/pickups", json={"id": shipment_id})

    async def track(self, tracking_no: str) -> dict:
        """追跡情報を取得する。"""
        res = httpx.get("https://api.carrier.example.com/v2/tracking")
        return res.json()
