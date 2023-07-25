package template

import "fmt"

var AccountActivation = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Account Activation</title>
    <link rel="stylesheet" href="activation.css">
	<style>
	body {
		font-family: Arial, sans-serif;
		background-color: #f0f0f0;
		margin: 0;
		padding: 0;
	}
	
	.container {
		max-width: 600px;
		margin: 30px auto;
		padding: 20px;
		background-color: #fff;
		border: 1px solid #ccc;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
		text-align: center;
	}
	
	h1 {
		color: #333;
	}
	
	.activation-btn {
		margin-top: 20px;
	}
	
	.activation-btn a {
		display: inline-block;
		padding: 10px 20px;
		background-color: #007bff;
		color: #fff;
		text-decoration: none;
		border-radius: 5px;
	}
	
	.activation-btn a:hover {
		background-color: #0056b3;
	}
	
	p {
		color: #666;
		line-height: 1.6;
	}
	
	</style>
</head>
<body>
    <div class="container">
        <h1>Account Activation</h1>
        <p>Dear User,</p>
        <p>Thank you for signing up with us. To activate your account, please click the button below:</p>
        <div class="activation-btn">
            <a href="http://localhost:5173/uas/activate/%s">Activate Account</a>
        </div>
        <p>If you did not sign up for an account on our website, please ignore this email.</p>
        <p>Best regards,</p>
        <p>Apollo Service</p>
    </div>
</body>
</html>
`

func ReturnHtmlTemplate(id string) string {
	return fmt.Sprintf(AccountActivation, id)
}
