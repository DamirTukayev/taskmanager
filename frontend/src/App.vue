<template>
  <v-card width="1000px">
    <v-tabs v-model="tab" bg-color="primary">
      <v-tab value="one">Задачи</v-tab>
      <v-tab value="two">Записи</v-tab>
    </v-tabs>

    <v-card-text >
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="one">
          <Tasks />
        </v-tabs-window-item>

        <v-tabs-window-item value="two"> <Entries /> </v-tabs-window-item>
      </v-tabs-window>
    </v-card-text>
  </v-card>
</template>
<script setup>
import { ref } from "vue";
import Tasks from "./components/Tasks.vue";
import Entries from "./components/Entries.vue";

const tab = ref(null);
</script>
<style scoped>
</style>
