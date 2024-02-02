<script lang="ts">
	export let name: string;
	export let label: string;
	export let required: boolean = false;
</script>

<div class="flex flex-col gap-1">
	<label for={name}>
		{label}
	</label>
	<input class="rounded-lg bg-slate-50 px-3 py-2" id={name} {name} type="password" {required} />
</div>
