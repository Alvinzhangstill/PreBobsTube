package web

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	if c.Ctx.Input.IsAjax() {
		//由于是ajax请求，因此地址是header里的Referer
		//returnURL := c.Ctx.Input.Refer()
		html := `
<strong class="popup-title">Login</strong>
<div class="popup-holder">
	<form action="https://bobs-tube.com/login/" data-form="ajax" method="post">
		<div class="generic-error hidden">
					</div>
		<div>
			<div class="row">
				<label for="login_username" class="field-label required">Username</label>
				<input type="text" name="username" id="login_username" class="textfield" placeholder="please enter login here"/>
				<div class="field-error down"></div>
			</div>

			<div class="row">
				<label for="login_pass" class="field-label required">Password</label>
				<input type="password" name="pass" id="login_pass" class="textfield"/>
				<div class="field-error down"></div>
			</div>

							<div class="row">
					<input type="checkbox" name="remember_me" id="login_remember_me" class="checkbox" value="1"/>
					<label for="login_remember_me">remember me</label>
				</div>

			<div class="bottom">
				<input type="hidden" name="action" value="login"/>
				<input type="hidden" name="email_link" value="https://bobs-tube.com/email/"/>
				<input type="submit" class="submit" value="Log in"/>
				<div class="links">
					<p><a href="https://bobs-tube.com/signup/" data-fancybox="ajax">Not a member yet? Sign up now for free!</a></p>
					<p>
						<a href="https://bobs-tube.com/reset-password/" data-fancybox="ajax">Forgot password?</a> /
						<a href="https://bobs-tube.com/resend-confirmation/" data-fancybox="ajax">Missing confirmation email?</a>
					</p>
				</div>
			</div>
		</div>
	</form>
</div>
`
		c.Ctx.WriteString(html)
	}

}
