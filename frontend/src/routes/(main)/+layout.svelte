<script lang="ts">
	import apiClient from '$lib/api/client';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	onMount(() => {
		if (data.user.token) {
			apiClient.setToken(data.user.token);
		}
	});
</script>

<slot />
