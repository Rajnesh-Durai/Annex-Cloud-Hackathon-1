<script>
  import { navigate } from 'svelte-routing';

  let fieldaLogo ="https://annexstorage.blob.core.windows.net/annexcontainer/fielda-logo1.png";
  let loginImage = "https://annexstorage.blob.core.windows.net/annexcontainer/Vector-Smart-Object.svg";
  
  let email = '';
  let password = '';
  let showPassword = false;
  let isPasswordInputFocused = false;
  let emailError = '';
  let passwordError = '';

  const handlePasswordToggle = () => {
    showPassword = !showPassword;
  };

  const handlePasswordInputFocus = () => {
    isPasswordInputFocused = true;
  };

  const handlePasswordInputBlur = () => {
    isPasswordInputFocused = false;
  };

  const validateEmail = () => {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  };

  const handleSubmit = async () => {
    console.log(email)
    console.log(password)
    if (!email) {
      emailError = 'Please enter your email';
    } else if (!validateEmail()) {
      emailError = 'Enter your email address';
    } else {
      emailError = '';
    }

    if (!password) {
      passwordError = 'Enter your password';
    } else {
      passwordError = '';
    }

    if (!emailError && !passwordError) {
      const loginuser = {
        email: email,
        password: password,
      };
      console.log(loginuser);

      await fetch(`http://localhost:8888/api/authentication/login`, {
        method: "POST",
        body: JSON.stringify(loginuser),
      })
        .then(response => {
          if (!response.ok) {
            throw new Error("Login failed");
          }
          return response.json();
        })
        .then(data => {
          console.log("Login successfully");
          console.log(data);
          //toast.success('LoggedIn Successfully');
          sessionStorage.setItem('token', data.data.token);
          navigate('/dashboard');
        })
        .catch(error => {
        //  toast.error("Invalid Credentials");
          console.error("login error in:", error);
        });
    }
  };
</script>
<div class="login-body">
  <div class="left">
      <img src={fieldaLogo} alt="Logo" id="fieldaLogo"/>
      <img src={loginImage} alt="Login" id="loginImage" />
      <h1>A complete solution for all your field operations.</h1>
  </div>
  <div class="right">
      <div class="login-form">
          <h1 id="loginText">Sign In</h1>
          <h3 id="subHeadline">Welcome back! Please enter email id and password</h3>
          <div class="form-group field">
              <input type="email" class="form-field" placeholder="" bind:value={email} required>
              <label for="email" class="form-label">Email Id</label>
              {#if emailError}
                  <div class="error-message">{emailError}</div>
              {/if}
          </div>
          <br />
          <div class="form-group field">
              {#if showPassword}
                      <input
                          type="text"
                          class="form-field"
                          placeholder=""
                          bind:value={password}
                          on:focus={handlePasswordInputFocus}
                          on:blur={handlePasswordInputBlur}
                          required
                      />
                  {:else}
                      <input
                          type="password"
                          class="form-field"
                          placeholder=""
                          bind:value={password}
                          on:focus={handlePasswordInputFocus}
                          on:blur={handlePasswordInputBlur}
                          required
                      />
                  {/if}
              
              <label for="password" class="form-label">Password</label>
              {#if passwordError}
                  <div class="error-message">{passwordError}</div>
              {/if}
          </div>
          <p id="forgotPwd">Forgot Password ?</p>
          <button class="login-button" on:click={handleSubmit}>Log In</button>
          <p id="registerPara"><i>Don't have an account ? </i>&nbsp;<b>Sign Up</b></p>
      </div>
  </div>
</div>
<style>

</style>