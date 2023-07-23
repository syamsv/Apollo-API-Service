<script>
	// @ts-nocheck
	import {
		Header,
		SkipToContent,
		Content,
		Theme
	} from 'carbon-components-svelte';
</script>
<Theme  persist persistKey="__carbon-theme" />

<Header company="Apollo" platformName="Web Scraper Proxy">
	<svelte:fragment slot="skip-to-content">
		<SkipToContent />
	</svelte:fragment>
</Header>

<Content>
	<slot />
</Content>
