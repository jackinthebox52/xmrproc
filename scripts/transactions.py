import requests

endpoint = 'http://localhost:8080/api/transactions/new'
data = {'amount': 10}

r = requests.post(endpoint, json=data)
print(r.json())