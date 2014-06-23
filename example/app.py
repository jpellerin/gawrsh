from bottle import post, request, route, run, template
import json
import redis


@route('/')
def index():
    return template('example/views/index')


@post('/api/poke')
def poke():
    send('poke', {'from': request.forms.get('user')})


@post('/api/say')
def say():
    send('say', {'from': request.forms.get('user'),
                 'message': request.forms.get('message')})


def send(topic, message):
    conn = redis.StrictRedis()
    conn.publish(
        'gawrsh-example', json.dumps({'topic': topic, 'message': message}))


if __name__ == '__main__':
    run(host='localhost', port=8000)
