<script>
  import { navigate } from 'svelte-routing';

  let fieldaLogo ="https://annexfe.blob.core.windows.net/fe-images/fielda-logo1.png";
  let loginImage = "https://annexfe.blob.core.windows.net/fe-images/Vector-Smart-Object.svg";
  
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
      emailError = 'Please enter a valid email address';
    } else {
      emailError = '';
    }

    if (!password) {
      passwordError = 'Please enter your password';
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
.login-body{
  display: flex;
  background-color: #ffffff;
  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
}
.left{
  width:130vh;
  height: 100vh;
  background-image: linear-gradient(to right,#3a5bff36, #2696ff54,#2696ff42);;
  border-radius:0 0px 160px 0 ;
}
.left h1{
  font-size: 32px;
  font-weight:700;
  margin-left: 130px;
  color: #1c1c1c;
  margin-top: 100px;
}
#fieldaLogo{
  margin-top: 40px;
  margin-left: 220px;
}
#loginImage{
  margin-top: 70px;
  animation-name: moving;
  animation-duration: 2s;
  animation-iteration-count: infinite;
  margin-left: 70px;
  width: 800px;
  height: 460px;
}
@keyframes moving {
  0% { transform: translate(0px, 0px)  }
    50% { transform: translate(0px, 12px) }
    100% { transform: translate(0px, 0px) }
}
.right{
  width: 100vh;
  height: 100vh;
  background-image: linear-gradient(to right,#ffffff, #ffffff,#ffffff,#ffffff);
}
.login-form{
  height: 65vh;
  max-width: 100%;
  margin-left: 120px;
  margin-top: 180px;  
}
#loginText{
  font-size: 32px;
  text-align: left;
  padding-top: 40px;
  color: #12a484;
}
#subHeadline{
  font-size: 20px;
  text-align: left;
  color: rgb(137, 135, 135);
  padding-bottom: 20px;

}
.form-group {
  position: relative;
  padding: 20px 0 0 20px;
  width: 100%;
  max-width: 280px;
  margin-left: 15px;
}

.form-field {
  width: 140%;
  border: none;
  border-bottom: 2px solid #000000;
  outline: 0;
  font-size: 18px;
  color: #2d2929c9;
  letter-spacing: 0rem;
  background: transparent;
  transition: border-color 0.2s;
  margin-bottom: 5px;
}

.form-field::placeholder {
  color: transparent;
}

.form-field:placeholder-shown ~ .form-label {
  font-size: 18px;
  cursor: text;
  top: 20px;
}

.form-label {
  position: absolute;
  top: 0;
  display: block;
  transition: 0.2s;
  font-size: 0px;
  color: #000000;
  pointer-events: none;
}

.form-field:focus {
  padding-bottom: 6px;
  padding-top: 8px;
  font-weight: 100;
  border-width: 3px;
  border-image: linear-gradient(to right, #343333, #1c1c1c);
  border-image-slice: 1;
}

.form-field:focus ~ .form-label {
  position: absolute;
  top: 0;
  display: block;
  transition: 0.2s;
  font-size: 22px;
  color: #151515;
  font-weight: 700;
}
.error-message{
  color: red;
}
.password-toggle{
  position: absolute;
  top: 20px;
  left: 280px;
  cursor: pointer;
}
#forgotPwd{
  margin-left: 256px;
  cursor: pointer;
  color: #343333;
  margin-top: 20px;
  font-size: 18px;
}
#forgotPwd:hover{
  font-weight: 500;
}
.login-button {
  text-transform: uppercase;
  background-color: #f3f7fe;
  color: #000000;
  font-weight: 600;
  outline: 2px solid green;
  border-radius: 8px;
  width: 180px;
  height: 40px;
  transition: .3s;
  cursor: pointer;
  margin-top: 50px;
  margin-bottom: 10px;
  margin-left: 120px;
  font-size: 18px;
}
.login-button:hover {
  background-color: #55abfb;
  box-shadow: 0 0 0 5px #3b83f65f;
  color: #fff;
}
#registerPara{
  font-size: 20px;
  text-align: left;
  padding-top: 20px;
  color: #343333;
  margin-top: 15px;
  margin-left: 60px;
}
#registerPara b{
  cursor: pointer;
  color: #12a484;
  font-size: 24px;
}
#registerPara b:hover{
  text-decoration: underline;
  color: #3A5CFF;
}
@media (max-width: 410px) and (min-width: 210px) {
  .left{
    display: none;
  }
  .right{
    background-image: url('../../assets/fielda-logo1.png');
    background-repeat: no-repeat;
    background-size: 220px;
    background-position: 17% 5%;
    
  }
  .login-form{
    height: 65vh;
    max-width: 100% ;
    margin-left: 30px;
    margin-top: 80px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
}
@media (max-width: 410px) and (min-width: 380px){
  .left{
    display: none;
  }
  .right{
    background-image: url('../../assets/fielda-logo1.png');
    background-repeat: no-repeat;
    background-size: 220px;
    background-position: 18% 5%;
    
  }
  .login-form{
    height: 65vh;
    max-width: 100% ;
    margin-left: 30px;
    margin-top: 80px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
}
@media (max-width: 539px) and (min-width: 381px){
  .left{
    display: none;
  }
  .right{
    background-image: url('../../assets/fielda-logo1.png');
    background-repeat: no-repeat;
    background-size: 220px;
    background-position: 16% 5%;
    
  }
  .login-form{
    height: 65vh;
    max-width: 100% ;
    margin-left: 30px;
    margin-top: 80px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
}
/* @media (max-width: 539px) and (min-width: 481px){
  .left{
    width: 38vh;
  }
  .left h1{
    font-size: larger;
  }
  #loginImage{
    height: 420px;
    width: 380px;
  }
  #fieldaLogo{
    height: 160px;
    width: 320px;
  }
  .login-form{
    height: 65vh;
    width: 50vh;
    margin-left: 20px;
    margin-top: 220px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
} */
@media (max-width: 767px) and (min-width: 540px){
  .left{
    display: none;
  }
  .right{
    background-image: url('../../assets/fielda-logo1.png');
    background-repeat: no-repeat;
    background-size: 220px;
    background-position: 16% 5%;
    
  }
  .login-form{
    height: 65vh;
    max-width: 100% ;
    margin-left: 20px;
    margin-top: 80px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
}
@media (max-width: 911px) and (min-width: 768px){
  .left{
    width: 38vh;
  }
  .left h1{
    font-size: larger;
  }
  #loginImage{
    height: 420px;
    width: 380px;
  }
  #fieldaLogo{
    height: 160px;
    width: 320px;
    margin-left: 30px;
  }
  .login-form{
    height: 65vh;
    width: 50vh;
    margin-left: 20px;
    margin-top: 220px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}

}
@media (max-width: 1023px) and (min-width: 912px){
  .left{
    width: 45vh;
  }
  .left h1{
    font-size: larger;
  }
  #loginImage{
    height: 520px;
    width: 480px;
  }
  #fieldaLogo{
    height: 160px;
    width: 320px;
    margin-left: 80px;
  }
  .login-form{
    height: 65vh;
    width: 50vh;
    margin-left: 40px;
    margin-top: 220px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}

}
@media (max-width: 1200px) and (min-width: 1024px){
  .left{
    display: none;
  }
  .right{
    background-image: url('../../assets/fielda-logo1.png');
    background-repeat: no-repeat;
    background-size: 220px;
    background-position: 16% 5%;
    
  }
  .login-form{
    height: 65vh;
    width: 100vh;
    margin-left: 70px;
    margin-top: 90px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}

}
@media (max-width: 1300px) and (min-width: 1201px){

.left{
    width: 70vh;
  }
  .left h1{
    font-size: 28px;
    margin-top: -50px;
  }
  #loginImage{
    height: 520px;
    width: 480px;
    margin-top: -80px;
    margin-left: 40px;
  }
  #fieldaLogo{
    height: 160px;
    width: 320px;
    margin-left: 120px;
  }
  .login-form{
    height: 65vh;
    width: 50vh;
    margin-left: 80px;
    margin-top: 150px;    
}
#loginText{
  padding-top: 40px;
}
#subHeadline{
  font-size: 12px;
  padding-bottom: 0px;
}
#forgotPwd{
  margin-left: 166px;
}
.login-button{
  margin-left: 70px;
}
}
</style>