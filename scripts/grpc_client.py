import asyncio

from grpclib.client import Channel

from ozonmp.srv_review_api.v1.srv_review_api_grpc import SrvReviewApiServiceStub
from ozonmp.srv_review_api.v1.srv_review_api_pb2 import DescribeReviewV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = SrvReviewApiServiceStub(channel)

        req = DescribeReviewV1Request(review_id=1)
        reply = await client.DescribeReviewV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
