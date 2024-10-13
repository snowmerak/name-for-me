<script lang="ts">
  import logo from './assets/images/logo-universal.png'
  import {Greet, GetModels, GetLanguages, Generate} from '../wailsjs/go/main/App.js'
  import {onMount} from "svelte";
  import markdown from "@wcj/markdown-to-html";
  import {Marked} from "@ts-stack/markdown";

  let models: Array<string> = []
  let selectedModel = 'qwen2.5-coder:1.5b'

  let languages = []
  let selectedLanguage = 'Go'

  let resultText = ''
  let summaryText = ''

  onMount(() => {
    GetModels().then((response) => {
      models = response
    }).catch((error) => {
      alert(error)
    })

    GetLanguages().then((response) => {
      languages = response
    }).catch((error) => {
      alert(error)
    })
  })

  $: console.log(selectedModel)
  $: console.log(selectedLanguage)
  $: console.log(summaryText)
</script>

<main>
  <form class="margin">
    <fieldset role="group">
      <select bind:value={selectedModel}>
        {#each models as model}
          <option value={model}>{model}</option>
        {/each}
      </select>
      <button on:click={() => {
        GetModels().then((response) => {
          models = response
        }).catch((error) => {
          alert(error)
        })
      }}>REFRESH</button>
    </fieldset>
  </form>
  <form class="margin">
    <fieldset role="group">
      <select bind:value={selectedLanguage}>
        {#each languages as language}
          <option value={language}>{language}</option>
        {/each}
      </select>
    </fieldset>
  </form>
  <div class="margin">
    <textarea on:change={(ev) => {
      summaryText = ev.target.value
    }} />
  </div>
  <button class="margin" on:click={() => {
    console.log(`selectedModel: ${selectedModel}, selectedLanguage: ${selectedLanguage}, summaryText: ${summaryText}`)
    Generate(selectedModel, selectedLanguage, summaryText).then((response) => {
      resultText = response
    }).catch((error) => {
      alert(error)
    })
  }}>
    GENERATE
  </button>
  <div class="margin">
    <article>
      {@html Marked.parse(resultText)}
    </article>
  </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .margin {
    margin: 12px 12px 0 12px;
  }
</style>