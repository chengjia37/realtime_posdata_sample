# coding: utf-8
import sys
import os
import boto3
from boto3.session import Session

REGION_NAME = "ap-northeast-1"
TABLE_NAME  = "REALTIME-POS-DATA-TEI"
profile = 'kinesis'

def main():
    session = Session(profile_name=profile)
    dynamodb = session.resource('dynamodb', region_name=REGION_NAME)
    table    = dynamodb.Table(TABLE_NAME)
    truncate_dynamo_items(table)

    return 0

def truncate_dynamo_items(dynamodb_table):

    # データ全件取得
    delete_items = []
    parameters   = {}
    while True:
        response = dynamodb_table.scan(**parameters)
        delete_items.extend(response["Items"])
        if ( "LastEvaluatedKey" in response ):
            parameters["ExclusiveStartKey"] = response["LastEvaluatedKey"]
        else:
            break

    # キー抽出
    key_names = [ x["AttributeName"] for x in dynamodb_table.key_schema ]
    delete_keys = [ { k:v for k,v in x.items() if k in key_names } for x in delete_items ]

    # データ削除
    with dynamodb_table.batch_writer() as batch:
        for key in delete_keys:
            batch.delete_item(Key = key)

    return 0


if __name__ == '__main__':
    ret = main()
    sys.exit(ret)
