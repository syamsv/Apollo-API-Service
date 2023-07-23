<script lang="ts">
	import { enhance } from '$app/forms';
	import { TextInput, PasswordInput, Button, Link, ClickableTile } from 'carbon-components-svelte';
	let email = '';
	let password = '';
	let firstname = '';
	let lastname = '';
	let disabled = true;

	$: disabled = !email || !password;
</script>

<svelte:head>
	<title>Sign in to Apollo Account</title>
</svelte:head>

<div class="signupContainer">
	<div class="loginContainer">
		<h4>Account information</h4>
		<br /><br />
		<form method="post" use:enhance>
			<TextInput
				required
				labelText="Email Address"
				name="email"
				placeholder="Enter your email address"
				type="email"
				bind:value={email}
			/>
			<div class="group-vertical">
				<TextInput
					bind:value={firstname}
					name="firstname"
					required
					labelText="First name"
					placeholder="Enter your first name..."
				/>
				<TextInput
					bind:value={lastname}
					required
					name="lastname"
					labelText="Last Name"
					placeholder="Enter your last name..."
				/>
			</div>
			<PasswordInput
				required
				name="password"
				type="password"
				labelText="Password"
				placeholder="Enter password..."
				bind:value={password}
			/>
			<Button type="submit" {disabled}
				>Create an Apollo account
			</Button>
		</form>
		<div class="vertically" />
		<ClickableTile href="/uas/login">Already have an account? Create one</ClickableTile>
	</div>
</div>

<style>
	.signupContainer {
		width: 100%;
		height: 100%;
	}

	.loginContainer {
		width: 50%;
		height: 100%;
	}

	.loginContainer form {
		display: flex;
		flex-direction: column;
		gap: 30px;
	}

	.group-vertical {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 10px;
	}
	@media screen and (max-width: 1200px) {
		.loginContainer {
			width: 100%;
            overflow: hidden;
        }
        
	}
</style>
