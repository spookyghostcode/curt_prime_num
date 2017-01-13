from flask import Flask, request, jsonify

app = Flask(__name__, static_url_path='')

@app.route('/primes', methods=["GET"])
def primes():
    
    if not request.args.get('max'):
        resp = jsonify({"Error": "Argument 'max' not provided'"})
        resp.status_code = 400
        return resp

    try:
        max_num = int(request.args.get('max'))
    except ValueError:
        resp = jsonify({"Error": "'max' must be an integer"})
        resp.status_code = 400
        return resp

    prime_list = []

    for i in range(2, (max_num+1)):
        for j in range(2, i):
            if i != j:
                if (i % j) == 0:
                    break
        else:
            prime_list.append(i)

    resp = jsonify({"Success": prime_list})
    return resp

@app.route('/', methods=["GET"])
def homepage():
    return app.send_static_file('index.html')

if __name__ == "__main__":
    app.run(port=80)
