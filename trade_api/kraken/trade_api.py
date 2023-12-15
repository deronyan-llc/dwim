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

def add_market_order(pair, type, qty):
    return private_post(
        'AddOrder', 
        dict (
            make_nonce(),
            **{
                "ordertype": "market",
                "type": type,
                "volume": qty,
                "pair": pair
            }
        )
    )

def list_open_orders():
    return private_post('OpenOrders', make_nonce())

def list_closed_orders():
    return private_post('ClosedOrders', make_nonce())

def query_orders(orders):
    return private_post(
        'QueryOrders', 
        dict (
            make_nonce(),
            **{
                "txid": orders,
                "trades": "true"
            }
        )
    )

def pp(dictionary):
    print(json.dumps(dictionary.json(), indent=4))

def __main__():
    print("Kraken trade api")
    #pp(public_get_with_pair("Depth", "XRPUSD"))
    #pp(public_get("Time"))
    #pp(public_get_with_pair("Ticker", "XBTUSD"))
    #pp(private_post('Balance', make_nonce()))
    #pp(add_market_order("XRPUSD", "buy", 10))
    #pp(get_asset_pairs("SOLOUSD"))
    #pp(add_market_order("XRPUSD", "sell", 10))
    #pp(list_open_orders())
    #pp(list_closed_orders())
    pp(query_orders("OWRSNA-LSZ7T-NIVGN5"))
    # pretty print dictionary
    #pp(query_orders("OWRSNA-LSZ7T-NIVGN5"))

__main__()
