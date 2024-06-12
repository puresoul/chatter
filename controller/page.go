// Package static serves static files like CSS, JavaScript, and images.
package controller


const page = `
<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>CodePen - Direct Messaging</title>
  <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,600" rel="stylesheet">

<meta name="viewport" content="width=device-width, initial-scale=1"><link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/meyer-reset/2.0/reset.min.css">
<link rel="stylesheet" href="/assets/style.css">

{{.Style}}

</head>
<body>
<!-- partial:index.partial.html -->
<div class="wrapper">
    <div class="container">
        <div class="left">
            <ul class="people">
                <li class="person" data-chat="person1">
                    <span class="name">NAME</span>
                    <span class="time">TIME</span>
                </li>
            </ul>
        </div>
        <div class="right">
            <div class="top"><span><span class="name">NAME</span></span></div>
            <div class="chat">
                <div class="conversation-start">
                    <span>TIME</span>
                </div>
                <div class="bubble you">
                    YOU
                </div>
                <div class="bubble me">
                    ME
                </div>
            </div>
            <div class="write">
                <input type="text" />
                <a href="javascript:;" class="write-link send"></a>
            </div>
        </div>
    </div>
</div>

</body>
</html>
`