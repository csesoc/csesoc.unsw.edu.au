<!--
  EventDisplay
  --
  This component contains a grid of event tiles containing basic information regarding each event.
  It is currently being used in the landing page.
  --
  Props:
    - events: list of event objects - containing the fields: id, name, start_time, fb_event_id, fb_cover_img
    - updated: the time since the JSON file of events was last updated.
-->

<template>
  <div id="event-display" class="content">
    <v-container data-cy="event-section" class="Container">
      <v-container>
        <HeaderTitle title="EVENTS" />
        <p class="text--md">
          {{ subtitle }}
        </p>
      </v-container>
      <!-- Facebook embedded event page viewer -->
      <v-container class="Embed">
        <iframe
          src="https://www.facebook.com/plugins/page.php?href=https%3A%2F%2Fwww.facebook.com%2Fcsesoc&tabs=timeline&width=340&height=500&small_header=false&adapt_container_width=true&hide_cover=false&show_facepile=true&appId"
          width="340"
          height="500"
          style="border:none;overflow:hidden"
          scrolling="no"
          frameborder="0"
          allowfullscreen="true"
          allow="autoplay; clipboard-write; encrypted-media; picture-in-picture; web-share"
        />
      </v-container>
      <!-- Display all events in a sliding component on desktop viewports. -->
      <!-- <v-slide-group data-cy="event-slider" show-arrows dark class="hidden-sm-and-down">
        <v-slide-item v-for="event in events" :key="event.id" style="margin-right: 20px;">
          <Event :event = "event"></Event>
        </v-slide-item>
      </v-slide-group> -->

      <!-- Otherwise, list all events on mobile. -->
      <!-- <v-row data-cy="event-list" class="hidden-md-and-up" v-for="event in events.slice(0,3)" :key="event.id">
          <Event :event = "event" style="margin: 0 auto; margin-bottom: 20px;"></Event>
      </v-row> -->
    </v-container>
  </div>
</template>

<script type="text/javascript">
// import Event from '@/components/Event';
import HeaderTitle from '@/components/HeaderTitle';

export default {
  name: 'EventDisplay',
  // Must be passed from parent object
  // items have title, image url (src), and link
  props: ['events', 'updated'],
  components: {
    // Event,
    HeaderTitle
  },
  data: () => ({
    subtitle:
      'We run a wide-variety of events for fun, learning new skills and careers. For full listings, check out the CSESoc Discord or our Facebook page!'
  })
};
</script>

<style scoped>
.Container {
  display: flex;
  flex-direction: row;
}
.Embed {
  display: flex;
  justify-content: center;
}
@media screen and (max-width: 700px) {
  .Container {
    flex-direction: column;
  }
  .Embed {
    margin-top: 50px;
  }
}
</style>
