<script lang="ts">
    export let itemName: string = "";
    export let updateTime: string = "";
    export let value: any = undefined;
    
	let selection: never[] = [];

    const dateNow = new Date().toJSON();
    let displayTime = updateTime.slice(0,10) + " " + updateTime.slice(11,16);
    if(dateNow.slice(0,4) != updateTime.slice(0,4)) {
        console.log("display full date");
        displayTime = updateTime.slice(0,10) + " " + updateTime.slice(11,16);
    } else if(dateNow.slice(5,7) != updateTime.slice(5,7)) {
        console.log("display month and day");
        displayTime = updateTime.slice(5,10) + " " + updateTime.slice(11,16);
    } else if(dateNow.slice(8,10) != updateTime.slice(8,10)) {
        console.log("display day");
        displayTime = updateTime.slice(8,10) + " " + updateTime.slice(11,16);

        // TODO: Add fancy functions to see 'X hours ago' or 'yesterday' or 'X days ago', maybe even update every minute?
        // TODO: Update timezone?
    }

    // setTimeout(() => {
    //     let checkboxElem = document.querySelector('.checkbox');
    //     // console.log(getComputedStyle(checkboxElem));
    // }, 3000);
</script>

<label class="wrapper">
    <input type="checkbox" bind:group={selection} value={value} class="checkbox" /> 
    {itemName}
    <span class="date">{displayTime}</span>
</label>

<style lang="scss">

    .wrapper {
        margin: var(--margin-small) auto;
        border: 1px solid #fafafa;
        border-radius: calc(var(--baseline)*0.5);
        padding: var(--padding);
        max-width: 30rem;

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