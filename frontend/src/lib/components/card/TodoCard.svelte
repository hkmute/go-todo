<script lang="ts">
	export let title: string;
	export let description = '';
</script>

<div
	class="flex-1 rounded-2xl bg-indigo-50/65 p-4 transition-colors duration-100 hover:bg-indigo-50/100"
>
	<div class="">{title}</div>
	<div class="py-1 text-sm">{description}</div>
</div>
