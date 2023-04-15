<script lang="ts">
    import { headerTitle } from "../stores";
    
</script>

<header>
    <div>
        <figure><a href="/">🥚</a></figure>
        <br/>
        <br/>
        <h1>{$headerTitle}</h1>
    </div>
</header>

<style lang="scss">

    header {
        position: absolute;
        top: 0;
        width: 100%;
        height: var(--header-height);

        background-image: url(https://images.unsplash.com/photo-1598965675045-45c5e72c7d05?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1171&q=80);
        background-size: cover;
        background-position: center;
        background-color: burlywood;
        
        div {
            height: 100%;
            padding: var(--padding-large);
            background-color: #0000007c;
            
            h1 {
                -webkit-text-stroke: 2px #0a0a0a;
                // text-stroke: 1px #0000007c;
            }

            figure {
                width: min-content;
                transition: 1s;
                a {
                    font-size: calc(var(--baseline)*3);
                    text-decoration: none;
                }
                &:hover {
                    transform: scale(1.2);
                }
            }
        }
    }

</style>