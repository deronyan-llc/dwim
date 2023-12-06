import urllib.parse
import hashlib
import hmac
import base64
import requests
import time
import json
import os

api_url = 'https://api.kraken.com'
api_key = os.environ.get('KRAKEN_API_KEY')
api_sec = os.environ.get('KRAKEN_API_SEC')

def make_nonce():
    return { "nonce":  str(int(time.time()*1000)) }

def get_kraken_signature(urlpath, data, secret):
    postdata = urllib.parse.urlencode(data)
    encoded = (str(data['nonce']) + postdata).encode()
    message = urlpath.encode() + hashlib.sha256(encoded).digest()

    mac = hmac.new(base64.b64decode(secret), message, hashlib.sha512)
    sigdigest = base64.b64encode(mac.digest())
    return sigdigest.decode()

# Attaches auth headers and returns results of a POST request
def kraken_request(uri_path, data):
    headers = {}
    headers['API-Key'] = api_key
    headers['API-Sign'] = get_kraken_signature(uri_path, data, api_sec)
    req = requests.post((api_url + uri_path), headers=headers, data=data)
    return req

def public_get(arg):
    resp = requests.get(api_url+'/0/public/'+arg)
    if resp.status_code != 200:
        raise requests.ApiError('StatusCode: {}'.format(resp.status_code))
    else:
        return resp

def public_get_with_pair(arg, pair):
    resp = requests.get(api_url+'/0/public/'+arg+'?pair='+pair)
    if resp.status_code != 200:
        raise requests.ApiError('StatusCode: {}'.format(resp.status_code))
    else:
        return resp
def private_post(arg, data):
    resp = kraken_request("/0/private/"+arg, data)
    if resp.status_code != 200:
        raise requests.ApiError('StatusCode: {}'.format(resp.status_code))
    else:
        return resp

def get_asset_pairs(pairs):
    return public_get("AssetPairs?pair="+pairs)

def add_market_order(pair, volume):
    return private_post(
        'AddOrder', 
        dict (
            make_nonce(),
            **{
                "ordertype": "market",
                "type": "buy",
                "volume": volume,
                "pair": pair
            }
        )
    )

def __main__():
    # Construct the request and print the result
    print("Kraken trade api")
    #print(public_get_with_pair("Depth", "XRPUSD").json())
    #print(public_get("Time").json())
    #print(public_get_with_pair("Ticker", "XBTUSD").json())
    print(private_post('Balance', make_nonce()).json())
    #print(add_market_order("XRPUSD", 10).json())
    #print(get_asset_pairs("SOLOUSD").json())
    # get key from env

__main__()
