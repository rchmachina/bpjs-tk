<template>
  <div
    v-if="isModalOpen"
    class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50"
  >
    <div
      class="bg-white rounded-lg shadow-lg max-w-2xl w-full max-h-96 overflow-y-auto"
    >
      <div class="p-5">
        <div class="flex justify-between items-center">
          <h3 class="text-lg font-bold">Child Details</h3>
          <button @click="closeModal" class="text-red-600 text-2xl">
            &times;
          </button>
        </div>

        <table class="table w-full mt-5">
          <thead>
            <tr>
              <th class="border px-4 py-2">ID</th>
              <th class="border px-4 py-2">Nominal</th>
              <th class="border px-4 py-2">Start Date</th>
              <th class="border px-4 py-2">End Date</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="child in childData" :key="child.id">
              <td class="border px-4 py-2">{{ child.id }}</td>
              <td class="border px-4 py-2">{{ child.nominal }}</td>
              <td class="border px-4 py-2">{{ child.startDate }}</td>
              <td class="border px-4 py-2">{{ child.endDate }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>

  <div class="p-20 ml-20">
    <div v-if="isWaiting" class="bg-green-100 text-green-800 p-4 rounded mb-4">
      Tunggu ya...(realtime bisa ditest via postman)
    </div>
    <div class="flex items-center">
      <button
        @click="addData"
        class="ml-2 bg-green-800 hover:text-green-800 text-white rounded-lg shadow-lg px-4 py-2"
      >
        Add Concurrent 1000 Data
      </button>
      <button
        @click="deleteData"
        class="ml-2 bg-red-800 hover:text-red-800 text-white rounded-lg shadow-lg px-4 py-2"
      >
        Clear
      </button>
    </div>
    <div class="bg-white card mt-4 border rounded-lg shadow-lg">
      <div class="overflow-x-auto max-h-96">
        <table class="table w-full">
          <thead>
            <tr>
              <th class="border px-4 py-2">ID</th>
              <th class="border px-4 py-2">Nominal</th>
              <th class="border px-4 py-2">Start Date</th>
              <th class="border px-4 py-2">End Date</th>
              <th class="border px-4 py-2">Detail</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in data" :key="item.id">
              <td class="border px-4 py-2">{{ item.id }}</td>
              <td class="border px-4 py-2">{{ item.nominal }}</td>
              <td class="border px-4 py-2">{{ item.startDate }}</td>
              <td class="border px-4 py-2">{{ item.endDate }}</td>
              <td class="border px-4 py-2">
                <button
                  @click="childDetail(item.id)"
                  class="ml-2 bg-blue-800 text-white px-4 py-1 rounded"
                >
                  Detail
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>
<script>
import { ref, onMounted } from "vue";
import axios from "axios";

export default {
  name: "DataView",
  setup() {
    const data = ref([]);
    const childData = ref([]);
    const isModalOpen = ref(false);
    const isWaiting = ref(false);

    const addData = () => {
      const link = `${import.meta.env.VITE_BACKEND_API_URL}/concurent-soal1`;
      isWaiting.value = true;
      axios
        .post(link)
        .then((response) => {
          console.log("Data posted successfully:", response.data);
          setTimeout(() => {
            isWaiting.value = false;
            fetchingGetAllData();
          }, 15000);
        })
        .catch((error) => {
          console.error("Error posting data:", error);
        });
    };

    const deleteData = () => {
      const link = `${import.meta.env.VITE_BACKEND_API_URL}/soal1`;
      isWaiting.value = true;
      axios
        .delete(link)
        .then((response) => {
          console.log("Data deleted successfully:", response.data);
          setTimeout(() => {
            isWaiting.value = false;
            fetchingGetAllData();
          }, 15000);
        })
        .catch((error) => {
          console.error("Error deleting data:", error);
        });
    };

    const fetchingGetAllData = () => {
      const link = `${
        import.meta.env.VITE_BACKEND_API_URL
      }/soal1?page=1&limit=1000`;
      axios
        .get(link)
        .then((response) => {
          data.value = response.data.data;
        })
        .catch((error) => {
          console.error(error);
        });
    };

    const childDetail = (id) => {
      const link = `${
        import.meta.env.VITE_BACKEND_API_URL
      }/child-soal1?parentId=${id}`;
      axios
        .get(link)
        .then((response) => {
          childData.value = response.data.data;
          isModalOpen.value = true; // Open the modal after getting child data
        })
        .catch((error) => {
          console.error(error);
        });
    };

    const closeModal = () => {
      isModalOpen.value = false; // Close the modal
    };

    onMounted(() => {
      fetchingGetAllData();
    });

    return {
      data,
      childData,
      isModalOpen,
      isWaiting,
      addData,
      deleteData,
      childDetail,
      closeModal,
    };
  },
};
</script>
