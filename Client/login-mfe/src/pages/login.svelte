<script>
	import { navigate } from 'svelte-routing';
	import '../css/global.css';
	let email = '';
	let password = '';
	let isPasswordInputFocused = false;
	let emailError = '';
	let passwordError = '';

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
		console.log(email);
		console.log(password);
		if (!email) {
			emailError = 'Enter your email';
		} else if (!validateEmail()) {
			emailError = 'Enter a valid email address';
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
				password: password
			};
			console.log(loginuser);

			await fetch(`http://localhost:8888/api/authentication/login`, {
				method: 'POST',
				body: JSON.stringify(loginuser)
			})
				.then((response) => {
					if (!response.ok) {
						throw new Error('Login failed');
					}
					return response.json();
				})
				.then((data) => {
					console.log('Login successfully');
					console.log(data);
					//toast.success('LoggedIn Successfully');
					sessionStorage.setItem('token', data.data.token);
					navigate('/dashboard');
				})
				.catch((error) => {
					//  toast.error("Invalid Credentials");
					console.error('login error in:', error);
				});
		}
	};
</script>

<div class="lg:flex">
	<div
		class="hidden lg:flex items-center justify-center bg-[url('https://www.annexcloud.com/wp-content/uploads/2023/07/Why-AC_banner-Desktop-1.jpg')] bg-cover p-4 flex-1 h-screen"
		style="border-top-right-radius: 450px;"
	>
		<div class="max-w-xs">
			<img
				src="https://www.annexcloud.com/wp-content/themes/berg-theme-child/dist/images/slider-accordion-bottom--bg.webp?v=271cf807fd7566fb54c7ba626ec5b8c0"
				alt=""
				class="relative top-full left-80 h-16 w-16 hidden xl:block"
			/>
		</div>
	</div>
	<div class="lg:w-10/12 xl:max-w-screen-sm">
		<div
			class="pt-8 pb-8 mt-20 bg-white-100 lg:bg-white flex justify-center lg:justify-center lg:px-10"
		>
			<div class="cursor-pointer flex items-center">
				<div>
					<img src="https://www.annexcloud.com/wp-content/uploads/2023/05/logo.svg" alt="Logo" />
				</div>
			</div>
		</div>
		<div class="mt-10 px-10 sm:px-24 md:px-48 lg:px-12 lg:mt-16 xl:px-24 xl:max-w-2xl">
			<h2
				class="text-center text-2xl text-indigo-900 font-display font-semibold lg:text-left xl:text-4xl
            xl:text-bold"
			>
				Login
			</h2>
			<div class="mt-12">
				<form>
					<div>
						<div class="text-sm font-bold text-gray-700 tracking-wide">Email Id</div>
						<input
							class="w-full text-lg py-2 border-b border-gray-300 focus:outline-none focus:border-indigo-500"
							type="email"
							bind:value={email}
						/>
						{#if emailError}
							<div class="error-message text-red-600">{emailError}</div>
						{/if}
					</div>
					<div class="mt-8">
						<div class="flex justify-between items-center">
							<div class="text-sm font-bold text-gray-700 tracking-wide">Password</div>
						</div>
						<input
							class="w-full text-lg py-2 border-b border-gray-300 focus:outline-none focus:border-indigo-500"
							type="password"
							bind:value={password}
							on:focus={handlePasswordInputFocus}
							on:blur={handlePasswordInputBlur}
						/>
						{#if passwordError}
							<div class="error-message text-red-600">{passwordError}</div>
						{/if}
					</div>
					<div class="mt-10">
						<button
							class="bg-indigo-500 text-gray-100 p-4 w-full rounded-full tracking-wide
                        font-semibold font-display focus:outline-none focus:shadow-outline hover:bg-indigo-700
                        shadow-lg"
							on:click={handleSubmit}
						>
							Login
						</button>
					</div>
				</form>
			</div>
		</div>
	</div>
</div>
