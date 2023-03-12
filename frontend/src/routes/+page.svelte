<script lang="ts">

</script>


<main>
    <ul class="lists">
        <li><a href="/list">Best bacon and egg</a></li>
        <li><a href="/list">Eggbread</a></li>
        <li><a href="/list">Groceries</a></li>
        <li><a href="/list">Ikea</a></li>
        <li><a href="/list">Lidl</a></li>
        <li><a href="/list">Party</a></li>
    </ul>

    <ul class="navigation">
        <li><a href="/account">Account</a></li>
        <li><a href="/settings">Settings</a></li>
        <li><a href="/about">About this site</a></li>
        <li><a href="/#">Log out</a></li>
    </ul>
</main>

<style lang="scss">

    main {
        padding: var(--padding);

        display: flex;
        justify-content: space-around;
        flex-direction: column;
    }

    ul {
        padding: 0;
        list-style: none;
        li {
            padding: var(--padding-small);
        }
    }

    a {
        font-size: calc(var(--baseline)*1.5);
        text-decoration: none;
        transition: 0.5s;
        &:hover {
            text-shadow: 0 0 5px #fafafa;
        }
    }

</style>