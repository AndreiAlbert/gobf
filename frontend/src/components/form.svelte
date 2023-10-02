<script>
    import CodeMirror from "svelte-codemirror-editor";
    /**
     * @type {string | null | undefined}
    */
    let value;
    async function handleRunClick() {
        console.log(value);
        const requestBody = {
            code: value
        }
        const response = await fetch('http://localhost:8080/processRequest', {
            method: 'POST', 
            headers: {
                'Content-type': 'application/json'
            },
            body: JSON.stringify(requestBody),
        })
        const jsonResponse = await response.json();
        console.log(jsonResponse);
    }
</script>

<div>
    <button on:click|preventDefault={handleRunClick} class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mt-1" type="submit">Run </button>
    <CodeMirror
    bind:value
    styles={{
        "&": {
            height: "30rem", 
            width: "100%"
        }
    }}
    />
</div>
