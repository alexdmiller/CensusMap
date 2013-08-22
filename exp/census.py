import httplib
import json
import urllib

def FIPS(code):
    print code
    return {
        'state': code[0:2],
        'county': code[2:5],
        'tract': code[5:11],
        'block-group': code[11]
        }

key = "ab995d23ecc36a2920db2262f9ea8a9003ab2098"
lat = str(47.598755)
long = str(-122.332764)

conn = httplib.HTTPConnection("data.fcc.gov")
conn.request("GET", "/api/block/find?format=json&latitude="+lat+"&longitude="+long+"&showall=true")
res = conn.getresponse()
resJSON = json.loads(res.read())
print resJSON

locationCodes = FIPS(resJSON['Block']['FIPS'])
print locationCodes

conn = httplib.HTTPConnection("api.census.gov")
requestParams = "/data/2011/acs5?get=B00001_001E&for=" +\
             "block+group:" + locationCodes['block-group'] +\
             "&in=state:" + locationCodes['state'] +\
             "+county:" + locationCodes['county'] +\
             "+tract:" + locationCodes['tract'] +\
             "&key=" + key
print requestParams
conn.request("GET", requestParams)
res = conn.getresponse()
print json.loads(res.read())
