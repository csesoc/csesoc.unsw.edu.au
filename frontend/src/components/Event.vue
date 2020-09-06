<!--
  EventDisplay
  --
  This component is an individual event tile used to represent an event.
  --
  Props:
    - event: an event - containing the fields: id, name, start_time, fb_event_id, fb_cover_img
-->

<template>
  <v-card class="event-card" :href="`https://facebook.com/${event.fb_event_id}`">
    <v-hover v-slot:default="{ hover }">
      <v-img :src="event.fb_cover_img">
        <v-expand-transition>
          <div
            v-if="hover"
            class="info-card transition-fast-in-fast-out v-card--reveal"
          >
            <strong><div id="name" class="text--lg" v-text="event.name"/></strong>
            <div class="event__place" v-if="event.place != undefined" v-text="'@ '+event.place"/>
            <div id="date" v-text="getDateString(event.start_time, event.end_time)"></div>
          </div>
        </v-expand-transition>
      </v-img>
    </v-hover>
  </v-card>
</template>

<script>
export default {
  name: 'Event',
  props: ['event'],
  methods: {
    getDateString(unix1, unix2) {
      const days = ['Sun', 'Mon', 'Tues', 'Wed', 'Thurs', 'Fri', 'Sat'];
      const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];
      const start = new Date(unix1 * 1000);
      const end = new Date(unix2 * 1000);

      const sDay = days[start.getDay()];
      const sDate = start.getDate();
      const sMonth = months[start.getMonth()];
      const sHour = (start.getHours() > 12) ? start.getHours() - 12 : start.getHours();
      const sMin = (start.getMinutes() === 0) ? '' : `:${(start.getMinutes() < 10 ? '0' : '')}${start.getMinutes()}`;
      const sMerid = (start.getHours() >= 12) ? 'PM' : 'AM';
      if (unix2 === undefined) {
        return `${sDay} ${sDate} ${sMonth} at ${sHour}${sMin} ${sMerid}`;
      }
      // if the date is the same, we just need to display start and end times of the event
      const eHour = (end.getHours() > 12) ? end.getHours() - 12 : end.getHours();
      const eMin = (end.getMinutes() === 0) ? '' : `:${(end.getMinutes() < 10 ? '0' : '')}${end.getMinutes()}`;
      const eMerid = (end.getHours() >= 12) ? 'PM' : 'AM';
      if (start.getDate() === end.getDate() && start.getMonth() === end.getMonth()) {
        // eg. Thurs 1 Jan 12 PM to 1:30 PM
        return `${sDay} ${sDate} ${sMonth} at ${sHour}${sMin} ${sMerid} - ${eHour}${eMin} ${eMerid}`;
      }

      // if the dates are different
      // eg. Thurs 1 Jan 12 PM to Thurs 2 Jan 1:30 PM
      const eDay = days[end.getDay()];
      const eDate = end.getDate();
      const eMonth = months[end.getMonth()];

      return `${sDay} ${sDate} ${sMonth} at ${sHour}${sMin} ${sMerid} - ${eDay} ${eDate} ${eMonth} at ${eHour}${eMin} ${eMerid}`;
    }
  }
};
</script>

<style scoped lang="scss">

  .event-card {
    width: 32em;
    max-width: 90vw;
  }

  .info-card {
    background-color: $primary-color;
    background: linear-gradient(0deg, $primary-color, $secondary-color-2);
    color: $light-color;
    box-sizing: border-box;
    height: 100%;
    padding-left: 30px;
    padding-right: 30px;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  #date {
    font-size: 1.3em;
  }

</style>
