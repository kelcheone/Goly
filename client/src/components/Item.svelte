<script>
  import Card from "./Card.svelte";
  import { Modals, closeModal, openModal } from "svelte-modals";
  import Modal from "./Modal.svelte";
  export let goly;
  let showCard = true;

  const update = async (data) => {
    const json = {
      redirect: data.redirect,
      goly: data.goly,
      random: data.random,
      id: data.id,
    };

    await fetch("http://localhost:8000/goly/", {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
  };

  const handleModalOpen = () => {
    openModal(Modal, {
      title: "Update Goly Link",
      send: update,
      goly: goly.goly,
      redirect: goly.redirect,
      random: goly.random,
      id: goly.id,
    });
  };
  async function deleteGoly() {
    if (
      confirm(
        "Are you sure you wish to delete this Goly link (" + goly.goly + ")?"
      )
    ) {
      await fetch("http://localhost:8000/goly/" + goly.id, {
        method: "DELETE",
      }).then((response) => {
        showCard = false;
        console.log(response);
      });
    }
  }
</script>

{#if showCard}
  <Card>
    <p>Goly http://localhost:8000/r/{goly.goly}</p>
    <p>Redirect: {goly.redirect}</p>
    <p>Clicked: {goly.clicked}</p>
    <button class="update" on:click={handleModalOpen(goly)}>Update</button>
    <button class="delete" on:click={deleteGoly}>Delete</button>
  </Card>
{/if}
<Modals>
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div slot="backdrop" class="backdrop" on:click={closeModal} />
</Modals>

<style>
  button {
    color: white;
    font-weight: bolder;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
    cursor: pointer;
  }
  button:active {
    background-color: #fff;
  }
  .update {
    background-color: yellowgreen;
  }
  .delete {
    background-color: red;
  }
</style>
