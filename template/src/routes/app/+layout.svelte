<script lang="ts">
	// @ts-nocheck
	import {
		Header,
		HeaderUtilities,
		SkipToContent,
		Content,
		HeaderAction,
		HeaderPanelDivider,
		HeaderPanelLinks,
		HeaderPanelLink,
		SideNav,
		SideNavItems,
		SideNavLink,
		SideNavMenu,
		SideNavMenuItem,
		Select,
		SelectItem,
		Theme
	} from 'carbon-components-svelte';
	import UserAvatarFilledAlt  from 'carbon-icons-svelte/lib/UserAvatarFilledAlt.svelte';
	import Dashboard from "carbon-icons-svelte/lib/Dashboard.svelte";
	import Api_1 from "carbon-icons-svelte/lib/Api_1.svelte";
	import ServerProxy from "carbon-icons-svelte/lib/ServerProxy.svelte";
	let isOpen = false;
	let isSideNavOpen = false;
	let theme = 'g90';
</script>

<Theme bind:theme persist persistKey="__carbon-theme" />

<Header company="Apollo" platformName="Gateways" bind:isSideNavOpen>
	<svelte:fragment slot="skip-to-content">
		<SkipToContent />
	</svelte:fragment>
	<HeaderUtilities>
		<HeaderAction bind:isOpen icon={UserAvatarFilledAlt} closeIcon={UserAvatarFilledAlt}>
			<HeaderPanelLinks>
				<HeaderPanelDivider>Themes</HeaderPanelDivider>
				<Select  bind:selected={theme}>
					<SelectItem value="white" text="White" />
					<SelectItem value="g10" text="Gray 10" />
					<SelectItem value="g80" text="Gray 80" />
					<SelectItem value="g90" text="Gray 90" />
					<SelectItem value="g100" text="Gray 100" />
				  </Select>
				<HeaderPanelDivider>Account actions</HeaderPanelDivider>
				<HeaderPanelLink>Logout</HeaderPanelLink>
			</HeaderPanelLinks>
		</HeaderAction>
	</HeaderUtilities>
</Header>
<SideNav bind:isOpen={isSideNavOpen} rail>
	<SideNavItems>
		<SideNavLink icon={Dashboard} text="Overview" />
		<SideNavMenu icon={Api_1} text="API Services">
			<SideNavLink icon={ServerProxy} href="/" text="Scraper Proxy" />
		</SideNavMenu>
	</SideNavItems>
</SideNav>
<Content>
	<slot />
</Content>
