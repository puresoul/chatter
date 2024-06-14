package controller

const style = `
<style>
@charset "UTF-8";
*, *:before, *:after {
  box-sizing: border-box;
}

:root {
  --white: #fff;
  --black: #000;
  --bg: #f8f8f8;
  --grey: #999;
  --dark: #1a1a1a;
  --light: #e6e6e6;
  --wrapper: 1000px;
  --blue: #00b0ff;
}

body {
  background-color: var(--bg);
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-rendering: optimizeLegibility;
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 400;
  background-image: url("https://s3-us-west-2.amazonaws.com/s.cdpn.io/382994/image.jpg");
  background-size: cover;
  background-repeat: none;
}

.wrapper {
  position: relative;
  left: 50%;
  width: var(--wrapper);
  height: 800px;
  transform: translate(-50%, 0);
}

.container {
  position: relative;
  top: 50%;
  left: 50%;
  width: 80%;
  height: 75%;
  background-color: var(--white);
  transform: translate(-50%, -50%);
}

.login {
  position: relative;
  background-color: var(--white);
transform: translate(50%, 50%);
    border-radius: 15px;
    box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
    padding: 10px 20px;
    transition: transform 0.2s;
    width: 500px;
    text-align: center;
}

label {
    display: block;
    width: 100%;
    margin-top: 10px;
    margin-bottom: 5px;
    text-align: left;
    color: #555;
    font-weight: bold;
}


input {
    display: block;
    width: 100%;
    margin-bottom: 15px;
    padding: 10px;
    box-sizing: border-box;
    border: 1px solid #ddd;
    border-radius: 5px;
}

button {
    padding: 15px;
    border-radius: 10px;
    margin-top: 15px;
    margin-bottom: 15px;
    border: none;
    color: white;
    cursor: pointer;
    background-color: #4CAF50;
    width: 100%;
    font-size: 16px;
}

.container .left {
  float: left;
  width: 37.6%;
  height: 100%;
  border: 1px solid var(--light);
  background-color: var(--white);
}
.container .left .top {
  position: relative;
  width: 100%;
  height: 96px;
  padding: 29px;
}
.container .left .top:after {
  position: absolute;
  bottom: 0;
  left: 50%;
  display: block;
  width: 80%;
  height: 1px;
  content: "";
  background-color: var(--light);
  transform: translate(-50%, 0);
}
.container .left input {
  float: left;
  width: 188px;
  height: 42px;
  padding: 0 15px;
  border: 1px solid var(--light);
  background-color: #eceff1;
  border-radius: 21px;
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 400;
}
.container .left input:focus {
  outline: none;
}
.container .left a.search {
  display: block;
  float: left;
  width: 42px;
  height: 42px;
  margin-left: 10px;
  border: 1px solid var(--light);
  background-color: var(--blue);
  background-image: url("https://s3-us-west-2.amazonaws.com/s.cdpn.io/382994/name-type.png");
  background-repeat: no-repeat;
  background-position: top 12px left 14px;
  border-radius: 50%;
}
.container .left .people {
  margin-left: -1px;
  border-right: 1px solid var(--light);
  border-left: 1px solid var(--light);
  width: calc(100% + 2px);
}
.container .left .people .person {
  position: relative;
  width: 100%;
  padding: 12px 10% 16px;
  cursor: pointer;
  background-color: var(--white);
}
.container .left .people .person:after {
  position: absolute;
  bottom: 0;
  left: 50%;
  display: block;
  width: 80%;
  height: 1px;
  content: "";
  background-color: var(--light);
  transform: translate(-50%, 0);
}
.container .left .people .person img {
  float: left;
  width: 40px;
  height: 40px;
  margin-right: 12px;
  border-radius: 50%;
  -o-object-fit: cover;
     object-fit: cover;
}
.container .left .people .person .name {
  font-size: 14px;
  line-height: 22px;
  color: var(--dark);
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 600;
}
.container .left .people .person .time {
  font-size: 14px;
  position: absolute;
  top: 16px;
  right: 10%;
  padding: 0 0 5px 5px;
  color: var(--grey);
  background-color: var(--white);
}
.container .left .people .person .preview {
  font-size: 14px;
  display: inline-block;
  overflow: hidden !important;
  width: 70%;
  white-space: nowrap;
  text-overflow: ellipsis;
  color: var(--grey);
}
.container .left .people .person.active, .container .left .people .person:hover {
  margin-top: -1px;
  margin-left: -1px;
  padding-top: 13px;
  border: 0;
  background-color: var(--blue);
  width: calc(100% + 2px);
  padding-left: calc(10% + 1px);
}
.container .left .people .person.active span, .container .left .people .person:hover span {
  color: var(--white);
  background: transparent;
}
.container .left .people .person.active:after, .container .left .people .person:hover:after {
  display: none;
}
.container .right {
  position: relative;
  float: left;
  width: 62.4%;
  height: 100%;
}
.container .right .top {
  width: 100%;
  height: 47px;
  padding: 15px 29px;
  background-color: #eceff1;
}
.container .right .top span {
  font-size: 15px;
  color: var(--grey);
}
.container .right .top span .name {
  color: var(--dark);
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 600;
}
.container .right .chat {
  position: relative;
  overflow: hidden;
  padding: 0 35px 92px;
  border-width: 1px 1px 1px 0;
  border-style: solid;
  border-color: var(--light);
  height: calc(100% - 48px);
  justify-content: flex-end;
  flex-direction: column;
}
.container .right .chat.active-chat {
  display: block;
  display: flex;
}
.container .right .chat.active-chat .bubble {
  transition-timing-function: cubic-bezier(0.4, -0.04, 1, 1);
}
.container .right .chat.active-chat .bubble:nth-of-type(1) {
  -webkit-animation-duration: 0.15s;
          animation-duration: 0.15s;
}
.container .right .chat.active-chat .bubble:nth-of-type(2) {
  -webkit-animation-duration: 0.3s;
          animation-duration: 0.3s;
}
.container .right .chat.active-chat .bubble:nth-of-type(3) {
  -webkit-animation-duration: 0.45s;
          animation-duration: 0.45s;
}
.container .right .chat.active-chat .bubble:nth-of-type(4) {
  -webkit-animation-duration: 0.6s;
          animation-duration: 0.6s;
}
.container .right .chat.active-chat .bubble:nth-of-type(5) {
  -webkit-animation-duration: 0.75s;
          animation-duration: 0.75s;
}
.container .right .chat.active-chat .bubble:nth-of-type(6) {
  -webkit-animation-duration: 0.9s;
          animation-duration: 0.9s;
}
.container .right .chat.active-chat .bubble:nth-of-type(7) {
  -webkit-animation-duration: 1.05s;
          animation-duration: 1.05s;
}
.container .right .chat.active-chat .bubble:nth-of-type(8) {
  -webkit-animation-duration: 1.2s;
          animation-duration: 1.2s;
}
.container .right .chat.active-chat .bubble:nth-of-type(9) {
  -webkit-animation-duration: 1.35s;
          animation-duration: 1.35s;
}
.container .right .chat.active-chat .bubble:nth-of-type(10) {
  -webkit-animation-duration: 1.5s;
          animation-duration: 1.5s;
}
.container .right .write {
  position: absolute;
  bottom: 29px;
  left: 30px;
  height: 42px;
  padding-left: 8px;
  border: 1px solid var(--light);
  background-color: #eceff1;
  width: calc(100% - 58px);
  border-radius: 5px;
}
.container .right .write input {
  font-size: 16px;
  float: left;
  width: 347px;
  height: 40px;
  padding: 0 10px;
  color: var(--dark);
  border: 0;
  outline: none;
  background-color: #eceff1;
  font-family: "Source Sans Pro", sans-serif;
  font-weight: 400;
}
.container .right .write .write-link.attach:before {
  display: inline-block;
  float: left;
  width: 20px;
  height: 42px;
  content: "";
  background-image: url("https://s3-us-west-2.amazonaws.com/s.cdpn.io/382994/attachment.png");
  background-repeat: no-repeat;
  background-position: center;
}
.container .right .write .write-link.smiley:before {
  display: inline-block;
  float: left;
  width: 20px;
  height: 42px;
  content: "";
  background-image: url("https://s3-us-west-2.amazonaws.com/s.cdpn.io/382994/smiley.png");
  background-repeat: no-repeat;
  background-position: center;
}
.container .right .write .write-link.send:before {
  display: inline-block;
  float: left;
  width: 20px;
  height: 42px;
  margin-left: 11px;
  content: "";
  background-image: url("https://s3-us-west-2.amazonaws.com/s.cdpn.io/382994/send.png");
  background-repeat: no-repeat;
  background-position: center;
}
.container .right .bubble {
  font-size: 16px;
  position: relative;
  display: inline-block;
  clear: both;
  margin-bottom: 8px;
  padding: 13px 14px;
  vertical-align: top;
  border-radius: 5px;
}
.container .right .bubble:before {
  position: absolute;
  top: 19px;
  display: block;
  width: 8px;
  height: 6px;
  content: " ";
  transform: rotate(29deg) skew(-35deg);
}
.container .right .bubble.you {
  float: left;
  color: var(--white);
  background-color: var(--blue);
  align-self: flex-start;
  -webkit-animation-name: slideFromLeft;
          animation-name: slideFromLeft;
}
.container .right .bubble.you:before {
  left: -3px;
  background-color: var(--blue);
}
.container .right .bubble.me {
  float: right;
  color: var(--dark);
  background-color: #eceff1;
  align-self: flex-end;
  -webkit-animation-name: slideFromRight;
          animation-name: slideFromRight;
}
.container .right .bubble.me:before {
  right: -3px;
  background-color: #eceff1;
}
.container .right .conversation-start {
  position: relative;
  width: 100%;
  margin-bottom: 27px;
  text-align: center;
}
.container .right .conversation-start span {
  font-size: 14px;
  display: inline-block;
  color: var(--grey);
}
.container .right .conversation-start span:before, .container .right .conversation-start span:after {
  position: absolute;
  top: 10px;
  display: inline-block;
  width: 30%;
  height: 1px;
  content: "";
  background-color: var(--light);
}
.container .right .conversation-start span:before {
  left: 0;
}
.container .right .conversation-start span:after {
  right: 0;
}

@keyframes slideFromLeft {
  0% {
    margin-left: -200px;
    opacity: 0;
  }
  100% {
    margin-left: 0;
    opacity: 1;
  }
}
@-webkit-keyframes slideFromLeft {
  0% {
    margin-left: -200px;
    opacity: 0;
  }
  100% {
    margin-left: 0;
    opacity: 1;
  }
}
@keyframes slideFromRight {
  0% {
    margin-right: -200px;
    opacity: 0;
  }
  100% {
    margin-right: 0;
    opacity: 1;
  }
}
@-webkit-keyframes slideFromRight {
  0% {
    margin-right: -200px;
    opacity: 0;
  }
  100% {
    margin-right: 0;
    opacity: 1;
  }
}
</style>
`