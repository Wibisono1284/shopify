{% if customer %}
  <p>Welcome back, {{ customer.first_name }}!</p>
  <a href="/account">Go to your account</a>
  <a href="/account/logout">Logout</a>
{% else %}
  <div class="login-form">
    <h2>Login</h2>
    <form method="post" action="/account/login">
      <input type="hidden" name="form_type" value="customer_login" />
      <input type="hidden" name="utf8" value="✓" />
      
      <label for="customer_email">Email</label>
      <input type="email" name="customer[email]" id="customer_email" autocorrect="off" autocapitalize="off" autofocus />
      
      <label for="customer_password">Password</label>
      <input type="password" name="customer[password]" id="customer_password" />
      
      <input type="submit" value="Login" />
    </form>
    <p><a href="/account/register">Create account</a></p>
    <p><a href="/account/recover">Forgot your password?</a></p>
  </div>
{% endif %}
