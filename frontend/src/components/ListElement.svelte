<script lang="ts">
    export let itemName = "";
    export let value: any = undefined;
    
	let selection: never[] = [];
    let date = new Date().getHours() + ':' + new Date().getMinutes();

    setTimeout(() => {
        let checkboxElem = document.querySelector('.checkbox');
        console.log(getComputedStyle(checkboxElem));
    }, 3000);
</script>

<label class="wrapper">
    <input type="checkbox" bind:group={selection} value={value} class="checkbox" /> 
    {itemName}
    <span class="date">{date}</span>
</label>

<style lang="scss">

    .wrapper {
        margin: var(--margin-small);
        border: 1px solid #fafafa;
        border-radius: calc(var(--baseline)*0.5);
        padding: var(--padding);

        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        gap: var(--baseline);

        .checkbox {
            // appearance: unset;
            height: calc(var(--baseline)*2.5);
            width: calc(var(--baseline)*2.5);
            border: 4px solid lightblue;
            border-radius: calc(var(--baseline)*0.5);
            & :checked {
                background-color: lightblue;
            }
        }

        .date {
            margin-left: auto;
        }
    }

</style>