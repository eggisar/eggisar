<script lang="ts">
    import ListElement from "../../../components/ListElement.svelte";
    import {page} from '$app/stores';
    import placeholderData from '../../../placeholderData.json';
    import Sorting from "../../../components/Sorting.svelte";
    import { headerTitle } from "../../../stores";

    const slug:string = $page.params.slug;
    let items: any[] = [];
    // console.log(slug);
    for(let list of placeholderData.list) {
        // console.log(list);
        // console.log(slug.includes(list.name));
        // console.log(slug.includes(list.id.toString()));

        if(slug.includes(list.name) && slug.includes(list.id.toString())){
            console.log("It's a match! ");
            headerTitle.set(list.name);
            for(let item of list.elements) {
                items.push(item);
            }
        }
    }


</script>


<main>
    <Sorting />

    {#if items.length > 0}
        {#each items as item}
            <ListElement itemName={item.name} updateTime={item.updated} />
        {/each}
    {:else} 
        <p>This list is empty.</p>
    {/if}
</main>


<style>

    main {
        padding: var(--padding);
    }

</style>