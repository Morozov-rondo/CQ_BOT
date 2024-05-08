package email_msg

const (
	Msg_ok = `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Заявка принята</title>
<style>
body {
font-family: Arial, sans-serif;
line-height: 1.5;
background-color: #f6f2f6;
margin: 0;
padding: 0;
color: #333;

}
.container {
max-width: 600px;
margin: 0 auto;
padding: 20px;
background: url('https://i.yapx.ru/XV6WY.jpg') no-repeat;
background-size:cover;
background-position: 0 100%;
border-radius: 15px;
}
h1 {
text-align: center;
}
p {
font-weight: 600;
}

.contacts{
margin-top: 20px;
border-left:3px solid #c1e3e9ef;

}
.contacts__item {
text-align: left;
padding: 0px 10px;
border-radius: 10px;
}
a {
color: #3b48f7;;
text-decoration: none;
}
span{
color: #333;
font-weight: 600;
}
</style>
</head>
<body>
<div class="container">
<h1>Здравствуйте, команда "%v"!</h1>
<p>Ваша заявка принята! Вы зарегистрированы на <b>%v – дата игры %v!</b></p>
<p>За день до игры мы отправим Вам сообщение в WhatsApp и уточним количество человек в команде.</p>
<p>До встречи на Сити Шоу!</p>
<p>С уважением, организаторы проекта "Сити Шоу"</p>

<p>Евгений Радибоженко, <a  href="tel:+79504305304">8-950-430-53-04</a><br>
Анна Радибоженко, <a  href="tel:+79029192892">8-902-919-28-92</a></p>

<div class="contacts">
<div class="contacts__item"><a href="http://vk.com/cityquizkrsk"><span>Сити Квиз:</span> vk.com/cityquizkrsk</a></div>
<div class="contacts__item"><a href="vk.com/aventurakrsk"><span>Авентура:</span> vk.com/aventurakrsk</a></div>
<div class="contacts__item"><a href="vk.com/sttwinskrsk"><span>Сейнт Твинс:</span> vk.com/sttwinskrsk</a></div>
<div class="contacts__item"><a href="flamp.ru/cityqkrsk"><span>Отзывы:</span> flamp.ru/cityqkrsk</a></div>
</div>
</div>
</body>
</html>`

	Msg_res = `<!DOCTYPE html>
<html lang="ru">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Заявка принята</title>
<style>
body {
font-family: Arial, sans-serif;
line-height: 1.5;
background-color: #f6f2f6;
margin: 0;
padding: 0;
color: #333;

}
.container {
max-width: 600px;
margin: 0 auto;
padding: 20px;
background: url('https://i.yapx.ru/XV6WY.jpg') no-repeat;
background-size:cover;
background-position: 0 100%;
border-radius: 15px;
}
h1 {
text-align: center;
}
p {
font-weight: 600;
}

.contacts{
margin-top: 20px;
border-left:3px solid #c1e3e9ef;

}
.contacts__item {
text-align: left;
padding: 0px 10px;
border-radius: 10px;
}
a {
color: #3b48f7;;
text-decoration: none;
}
span{
color: #333;
font-weight: 600;
}
</style>
</head>
<body>
<div class="container">
<h1>Здравствуйте, команда "%v"!</h1>
<p>Вы подали заявку на <b>%v – дата игры %v!</b></p>
<p>К сожалению, на данный момент все столы заняты. Если место освободится (а такое бывает), мы с вами свяжемся!</p>
<p>С уважением, организаторы проекта "Сити Шоу"</p>

<p>Евгений Радибоженко, <a  href="tel:+79504305304">8-950-430-53-04</a><br>
Анна Радибоженко, <a  href="tel:+79029192892">8-902-919-28-92</a></p>

<div class="contacts">
<div class="contacts__item"><a href="http://vk.com/cityquizkrsk"><span>Сити Квиз:</span> vk.com/cityquizkrsk</a></div>
<div class="contacts__item"><a href="vk.com/aventurakrsk"><span>Авентура:</span> vk.com/aventurakrsk</a></div>
<div class="contacts__item"><a href="vk.com/sttwinskrsk"><span>Сейнт Твинс:</span> vk.com/sttwinskrsk</a></div>
<div class="contacts__item"><a href="flamp.ru/cityqkrsk"><span>Отзывы:</span> flamp.ru/cityqkrsk</a></div>
</div>
</div>
</body>
</html>`
)
