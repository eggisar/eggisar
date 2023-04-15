<!-- code and elements are always visible in here -->

<script lang="ts">
    import Header from "../components/Header.svelte";
    import { headerTitle } from "../stores";
    headerTitle.set("My lists");
</script>

<Header />
<!-- <main> -->
    <slot></slot>
<!-- </main> -->

<svelte:head>
    <link rel="preconnect" href="https://fonts.googleapis.com">
    <!-- <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin> -->
    <link href="https://fonts.googleapis.com/css2?family=Ubuntu:ital,wght@0,300;0,400;0,500;0,700;1,300;1,400;1,500;1,700&display=swap" rel="stylesheet">
</svelte:head>

<style lang="scss">
    :global(main) {
        margin-top: var(--header-height);
        width: 100dvw;
        min-height: calc(100dvh - var(--header-height));
        background-color: #0a0a0a;
    }
    
    :global(*) {
        margin: 0;
        box-sizing: border-box;

        --baseline: 1rem;

        --header-height: calc(var(--baseline)*11);

        --padding: var(--baseline);
        --padding-small: calc(var(--baseline)*0.5);
        --padding-large: calc(var(--baseline)*1.5);

        --margin: var(--baseline);
        --margin-small: calc(var(--baseline)*0.5);
        --margin-large: calc(var(--baseline)*1.5);

        font-family: 'Ubuntu', sans-serif;
        color: #fafafa;
    }

</style>