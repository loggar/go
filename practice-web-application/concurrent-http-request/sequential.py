import requests

print(len(requests.get('http://localhost:8080/param1').text))
print(len(requests.get('http://localhost:8080/param2').text))
print(len(requests.get('http://localhost:8080/param3').text))

## python sequential.py
