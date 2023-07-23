<script lang="ts">
	import { enhance } from "$app/forms";
     import { PasswordInput, TextInput, Button,ClickableTile, Link } from "carbon-components-svelte";
     let email = '';
	let password = '';
	let disabled = true;
     $: disabled = !email || !password;
</script>
<div class="Container">
    <div class="halfContainer">
        <h4>Welcome Back</h4>
        <br><br>
        <form method="post" use:enhance>
            <TextInput
            required
            labelText="Email Address"
            placeholder="Enter your email address"
            name="email"
            type="email"
            bind:value={email}
        />
        <PasswordInput
            required
            type="password"
            labelText="Password"
            name="password"
            placeholder="Enter password..."
            bind:value={password}
        />
        <div class="forgot-password">
            <Link inline href="/uas/reset-password">Forgot Password</Link>
        </div>
        <Button type="submit" {disabled}
				>Login to Apollo
			</Button>
        </form>
        <div class="vertically" />
		<ClickableTile href="/uas/register">Dont have an account? Create one</ClickableTile>
    </div>
</div>

<style>
    .Container{
        width: 100%;
        height: 100%;
    }

    .halfContainer{
        width: 50%;
        height: 100%;
    }

    .halfContainer form{
        display: flex;
		flex-direction: column;
		gap: 30px;
    }
    .forgot-password {
		margin-top: 5px;
		width: 100%;
		display: flex;
		justify-content: end;
	}

    @media screen and (max-width: 1200px) {
		.halfContainer {
			width: 100%;
            overflow: hidden;
		}
	}
</style>