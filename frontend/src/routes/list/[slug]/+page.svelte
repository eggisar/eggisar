<script lang="ts">
    import Header from "../../../components/Header.svelte";
    import ListElement from "../../../components/ListElement.svelte";
    import {page} from '$app/stores';
    import placeholderData from '../../../placeholderData.json';
    
    const slug:string = $page.params.slug;
    let listName = "name of list";
    let items: any[] = [];
    // console.log(slug);
    for(let list of placeholderData.list) {
        // console.log(list);
        // console.log(slug.includes(list.name));
        // console.log(slug.includes(list.id.toString()));

        if(slug.includes(list.name) && slug.includes(list.id.toString())){
            console.log("It's a match! ");
            listName = list.name;
            for(let item of list.elements) {
                items.push(item);
            }
        }
    }


</script>

<Header title={listName} />
<main>
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