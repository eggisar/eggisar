<script lang="ts">
    import Header from "../../components/Header.svelte";
    import ListElement from "../../components/ListElement.svelte";

    const listName = "name of list";

</script>

<Header title={listName} />
<main>
    <ListElement itemName="eggs" />
    <ListElement itemName="milk" />
    <ListElement itemName="flour" />
</main>

<style>

    main {
        padding: var(--padding);
    }

</style>