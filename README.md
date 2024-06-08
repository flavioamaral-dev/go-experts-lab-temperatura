# go-experts-lab-temperatura
Sistema de temperatura por CEP


## Requisitos:

- O sistema deve receber um CEP válido de 8 digitos
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
- Em caso de sucesso:
Código HTTP: 200
Response Body: { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
- Em caso de falha, caso o CEP não seja válido (com formato correto):
Código HTTP: 422
Mensagem: invalid zipcode
​​​- Em caso de falha, caso o CEP não seja encontrado:
Código HTTP: 404
Mensagem: can not find zipcode
Deverá ser realizado o deploy no Google Cloud Run.

## Rodando em Google Cloud Run:
https://go-experts-lab-temperatura-5y6su7iqaq-uc.a.run.app/api/temperatura?cep=01002020
