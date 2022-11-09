<script>
  import { Modals, closeModal, openModal } from "svelte-modals";
  import Modal from "./Modal.svelte";

  const createGoly = async (data) => {
    const json = {
      redirect: data.redirect,
      goly: data.goly,
      random: data.random,
    };
    await fetch("http://localhost:8000/goly/", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(json),
    }).then((res) => {
      console.log(res);
    });
  };

  const handleOpen = () => {
    openModal(Modal, {
      title: "Create New Goly link",
      send: createGoly,
      redirect: "",
      random: false,
    });
  };
</script>

<button on:click={handleOpen}>New</button>

<style>
  button {
    background-color: green;
    color: white;
    font-weight: bold;
    border: none;
    padding: 0.75rem;
    border-radius: 4px;
    cursor: pointer;
  }

  button:is(:hover, :active) {
    background-color: greenyellow;
    color: black;
  }
</style>
