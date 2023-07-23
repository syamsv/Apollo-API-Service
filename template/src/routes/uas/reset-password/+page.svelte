<script lang="ts">
	import { enhance } from '$app/forms';
	import { TextInput, Button, ClickableTile } from 'carbon-components-svelte';
	let email = '';
	let disabled = true;

	$: disabled = !email;
</script>

<svelte:head>
	<title>Reset your Apollo Account Password</title>
</svelte:head>

<div class="signupContainer">
	<div class="loginContainer">
		<h4>Reset Your Password</h4>
		<br /><br />
		<form method="post" use:enhance>
			<TextInput
				required
				labelText="Email Address"
				placeholder="Enter your email address"
                name="email"
				type="email"
                helperText="If there is an account associated with given mail, you will receive an email with instructions to reset the password"
				bind:value={email}
			/>
			<Button type="submit" {disabled}>Request Password Reset</Button>
		</form>
		<div class="vertically" />
		<ClickableTile href="/uas/login">Already have an account? Login</ClickableTile>
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

	@media screen and (max-width: 1200px) {
		.loginContainer {
			width: 100%;
            overflow: hidden;
		}
	}
</style>
