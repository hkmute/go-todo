<script lang="ts">
	import { enhance } from '$app/forms';
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';
	import Button from '$lib/components/button/Button.svelte';
	import Input from '$lib/components/input/Input.svelte';
	import PasswordInput from '$lib/components/input/PasswordInput.svelte';
	import type { ActionData } from './$types';

	export let form: ActionData;


</script>

<div class="flex h-dvh items-center justify-center bg-indigo-50 p-4">
	<div class="flex w-[400px] flex-col rounded-2xl bg-indigo-100 p-8 text-indigo-900">
		<h1 class="pb-10 text-center text-3xl font-semibold">Register</h1>
		<form class="flex flex-col gap-2" method="POST" use:enhance={() => {
			return async ({ result, update }) => {
				if (result.type === 'redirect') {
					alert('Registration successful!');
				}
				update();
			};
		}}>
			<Input label="Username" name="username" />
			<PasswordInput label="Password" name="password" required />
			<PasswordInput label="Confirm Password" name="confirm-password" required />
			<div class="whitespace-pre-wrap">
				{#if form?.message}
					<ErrorMessage>
						{form.message}
					</ErrorMessage>
				{:else}
					{' '}
				{/if}
			</div>
			<div class="flex w-full flex-col gap-2 pt-4">
				<Button label="Register" type="submit" />
				<a href="/login">
					<Button variant="text" label="Already have an account?" />
				</a>
			</div>
		</form>
	</div>
</div>
