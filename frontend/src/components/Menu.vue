<!--
  Menu
  --
  This component contains an overlay with links to all the pages of the website.
  The links have a grow-on-hover feature.
-->

<template>
  <div @click.stop="hide" class="dimmed no-scroll">
    <div class="frame">
      <!-- Logo -->
      <RouterLink to="/" v-on:click.native="showMenu = false">
        <div class="fill-height" style="max-height: 64px; max-width:100px">
          <v-img src="@/assets/csesoc-logo-white.svg" />
        </div>
      </RouterLink>
      <div class="spacing-md" />
      <div class="row">
        <!-- Page Links -->
        <div class="page-links-container">
          <RouterLink to="/about" class="link">
            <h2>001 | About</h2>
          </RouterLink>
          <!-- <RouterLink to="/" class="link">
            <h2>Events</h2>
          </RouterLink> -->
          <RouterLink to="/resources" class="link">
            <h2>010 | Resources</h2>
          </RouterLink>
          <RouterLink to="/sponsors" class="link">
            <h2>011 | Sponsors</h2>
          </RouterLink>
          <RouterLink to="/engage" class="link">
            <h2>100 | Engage</h2>
          </RouterLink>
        </div>
        <!-- Social Links -->
        <div class="social-links-container">
          <label class="link">
            <p>Facebook</p>
          </label>
          <label class="link">
            <p>Twitter</p>
          </label>
          <label class="link">
            <p>Instagram</p>
          </label>
          <label class="link">
            <p>Youtube</p>
          </label>
          <label class="link">
            <p>LinkedIn</p>
          </label>
          <label class="link">
            <p>Discord Community</p>
          </label>
          <label class="link">
            <p>Slack Community</p>
          </label>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Menu',
  methods: {
    hide() {
      // Relaying status to parent
      this.$emit('shown', false);
    }
  }
};
</script>

<style lang="scss" scoped>
.dimmed {
  z-index: 5; /* display on top of nav-bar */
  display: block;
  position: fixed;
  height: 100%;
  width: 100%;
  top: 0;
  left: 0;
  right: 0;
  background: $dark-color-1;
}

.spacing-md {
  height: $space-md;
}

.spacing-xs {
  height: $space-xs;
}

.spacing-xxs {
  height: $space-xxs;
}

.spacing-xxxs {
  height: $space-xxxs;
}

.frame {
  margin-top: $space-xxs;
  margin-right: 4.5em;
  margin-bottom: $space-xs;
  margin-left: 4.5em;
}

.page-links-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 50%;

  .link {
    text-decoration: none; /* Remove underline from links */
    color: white;
    padding: $space-xs/2; /* Add some padding */
    padding-left: 0px; /* Keep it left aligned */
    cursor: pointer; /* To simulate a link */
    font-size: $text-md;
    font-weight: normal;
    transition: all .2s ease-in-out;
    width: 500px;

    h2 {
      cursor: pointer; /* To simulate a link */
    }
  }

  .link:hover {
    color: white;
    font-weight: bolder;
    font-size: $text-lg;
  }
}

.social-links-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  width: 50%;

  .link {
    text-decoration: none; /* Remove underline from links */
    text-align: right;
    color: white;
    padding: $space-xxxs/2; /* Add some padding */
    padding-right: 0px; /* Keep it right aligned */
    cursor: pointer; /* To simulate a link */
    // font-size: $text-md;
    font-weight: normal;
    transition: all .2s ease-in-out;
    width: 200px;

    p {
      cursor: pointer; /* To simulate a link */
    }
  }

  .link:hover {
    color: white;
    font-weight: bolder;
    font-size: $text-sm;
  }
}

/* Create two equal columns that floats next to each other */
.column {
  float: left;
  width: 50%;
  padding: 10px;
  height: 300px; /* Should be removed. Only for demonstration */
}

/* Clear floats after the columns */
.row:after {
  content: "";
  display: table;
  clear: both;
  height: 1000px;
}
</style>
