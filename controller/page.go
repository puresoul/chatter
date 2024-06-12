// Package static serves static files like CSS, JavaScript, and images.
package controller


const page = `
<!DOCTYPE html>
<html lang="en" >
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/meyer-reset/2.0/reset.min.css">
{{.}}

</head>

<body>
<div class="wrapper" id="chat">
</div>

</body>

<script>
let user = "";

function setName(e) {
user = e.id;
}

setInterval(function(){
   $( '#chat' ).load('/chat/?username='+user+' #chat');
}, 5000);
</script>
<script src="https://code.jquery.com/jquery-3.7.0.js"></script>
</html>
`