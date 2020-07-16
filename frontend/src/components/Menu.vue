<template>
  <div class="dimmed">
    <div class="vertical-menu" >
      <a @click.stop="showMenu = !showMenu" href="/#/about">About</a>
      <a href="#">Shop</a>
      <a href="#">Events</a>
      <a href="#">CSE Resources</a>
      <a href="/#/engage">Engage</a>
      <a href="#">Sponsors</a>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      
    };
  }
};
</script>

<style>
.dimmed {
  z-index: 2;
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  right: 0;
  background: rgba(0, 0, 0, 0.9);
  font-family: "Quicksand", sans-serif;
}

.vertical-menu {
  width: 500px; /* Set a width if you like */
  margin: auto;
  padding-top: 200px;
}

.vertical-menu a {
  color: #1976D2;
  display: block; /* Make the links appear below each other */
  padding: 12px; /* Add some padding */
  text-decoration: none; /* Remove underline from links */
  text-align: center;
  font-size: 35px;
  font-weight: bold;
}

.vertical-menu a:hover {
  color: white;
}
</style>
