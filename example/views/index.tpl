<!doctype html>
<html class="no-js" lang="">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <title></title>
        <meta name="description" content="">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <!-- Place favicon.ico and apple-touch-icon(s) in the root directory -->
        <link href='http://fonts.googleapis.com/css?family=Francois+One|Merriweather' rel='stylesheet' type='text/css'>
        <link rel="stylesheet" href="/static/css/normalize.css">
        <link rel="stylesheet" href="/static/css/main.css">
    </head>
    <body>
        <!--[if lt IE 8]>
            <p class="browsehappy">You are using an <strong>outdated</strong> browser. Please <a href="http://browsehappy.com/">upgrade your browser</a> to improve your experience.</p>
        <![endif]-->

        <!-- Add your site or application content here -->
        <h1>Hello!</h1> 

        <p>This is Gawrsh's example page, made using HTML5 Boilerplate.</p>

        <div class="col">
        <label>Who are you?</label>
        <input type="text" name="user" id="user">

        <hr>
        
        <button id="poke">Poke everyone</button>

        <label>Or say something instead</label>
        <textarea name="message" id="message"></textarea>
        
        <button id="say">Say it!</button>
        </div>

        <div class="col">
        <ul id="events"></ul>
        </div>

        <script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
        <script src="/static/js/main.js"></script>
    </body>
</html>
