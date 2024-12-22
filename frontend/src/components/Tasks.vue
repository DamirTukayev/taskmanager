<template>
  <v-container>
    <!-- Кнопка для открытия диалога -->
    <v-btn color="primary" @click="dialog = true">Добавить задачу</v-btn>

    <!-- Диалог с формой для добавления или редактирования задачи -->
    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title class="headline">{{ isEditing ? 'Редактировать задачу' : 'Добавить новую задачу' }}</v-card-title>
        <v-card-text>
          <v-form @submit.prevent="isEditing ? updateTask() : createTask">
            <v-text-field
              v-model="newTask.description"
              label="Описание"
              required
            ></v-text-field>
            <v-text-field
              v-model="newTask.time"
              label="Время (например: 12:30)"
              required
            ></v-text-field>
            <v-text-field
              v-model="newTask.date"
              label="Дата (например: 2024-12-10)"
              required
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions>
          <v-btn text @click="dialog = false">Отменить</v-btn>
          <v-btn color="primary" @click="isEditing ? updateTask() : createTask()">{{ isEditing ? 'Изменить задачу' : 'Добавить задачу' }}</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <hr class="my-7" />

    <!-- Список задач -->
    <v-row>
      <v-col v-for="task in tasks" :key="task.id" cols="12" sm="6" md="4">
        <v-card>
          <v-card-title class="d-flex justify-space-between align-center">
            <v-checkbox
              v-model="task.status"
              :label="'Запись ' + task.id"
              @change="toggleTaskStatus(task)"
              hide-details
            ></v-checkbox>
            <v-menu offset-y>
              <template v-slot:activator="{ props }">
                <v-btn variant="text" icon="mdi-menu" v-bind="props"></v-btn>
              </template>
              <v-list>
                <v-list-item @click="editTask(task)">
                  <v-list-item-title>Изменить</v-list-item-title>
                </v-list-item>
                <v-list-item @click="deleteTask(task.id)">
                  <v-list-item-title>Удалить</v-list-item-title>
                </v-list-item>
              </v-list>
            </v-menu>
          </v-card-title>
          <v-card-subtitle>{{ task.description }}</v-card-subtitle>
          <v-card-text>
            <div>Время: {{ task.time }}</div>
            <div>Дата: {{ task.date }}</div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      tasks: [],
      newTask: {
        description: "",
        time: "",
        date: "",
        status: false,
      },
      dialog: false,
      isEditing: false, // Флаг для определения, редактируем ли мы задачу
      currentTaskId: null, // Хранит ID задачи для редактирования
    };
  },
  mounted() {
    this.getAllTasks();
  },
  methods: {
    async getAllTasks() {
      const url = "http://localhost:8080/tasks";
      try {
        const response = await fetch(url);
        if (!response.ok) {
          throw new Error("Ошибка при загрузке данных");
        }
        const data = await response.json();
        this.tasks = data.sort((a, b) => b.status - a.status);
      } catch (error) {
        console.error("Ошибка при загрузке задач:", error);
      }
    },

    async createTask() {
      const url = "http://localhost:8080/tasks";
      try {
        const response = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(this.newTask),
        });

        if (!response.ok) {
          throw new Error("Ошибка при создании задачи");
        }

        const data = await response.json();
        this.tasks.push(data); // Добавляем задачу в список
        this.resetNewTask(); // Сброс формы
        this.dialog = false; // Закрываем диалог
      } catch (error) {
        console.error("Ошибка при создании задачи:", error);
      }
    },

    async updateTask() {
      const url = `http://localhost:8080/tasks/${this.currentTaskId}`;
      try {
        const response = await fetch(url, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(this.newTask),
        });

        if (!response.ok) {
          throw new Error("Ошибка при обновлении задачи");
        }

        const data = await response.json();
        const index = this.tasks.findIndex((task) => task.id === this.currentTaskId);
        if (index !== -1) {
          this.tasks[index] = data; // Обновляем задачу в списке
        }
        this.resetNewTask(); // Сброс формы
        this.dialog = false; // Закрываем диалог
      } catch (error) {
        console.error("Ошибка при изменении задачи:", error);
      }
    },

    async deleteTask(taskId) {
      const url = `http://localhost:8080/tasks/${taskId}`;
      try {
        const response = await fetch(url, {
          method: "DELETE",
        });

        if (!response.ok) {
          throw new Error("Ошибка при удалении задачи");
        }

        this.tasks = this.tasks.filter((task) => task.id !== taskId); // Удаляем задачу из списка
      } catch (error) {
        console.error("Ошибка при удалении задачи:", error);
      }
    },

    resetNewTask() {
      this.newTask = {
        description: "",
        time: "",
        date: "",
        status: false,
      };
      this.isEditing = false;
      this.currentTaskId = null;
    },

    editTask(task) {
      this.newTask = { ...task }; // Заполняем поля формы данными задачи
      this.isEditing = true; // Устанавливаем флаг редактирования
      this.currentTaskId = task.id; // Запоминаем ID редактируемой задачи
      this.dialog = true; // Открываем диалог
    },

    async toggleTaskStatus(task) {
      const url = `http://localhost:8080/tasks/${task.id}`;
      const updatedTask = {
        ...task,
        status: task.status,
      };

      try {
        const response = await fetch(url, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(updatedTask),
        });

        if (!response.ok) {
          throw new Error("Ошибка при обновлении задачи");
        }

        const data = await response.json();
        const index = this.tasks.findIndex((t) => t.id === task.id);
        if (index !== -1) {
          this.tasks[index] = data;
        }
      } catch (error) {
        console.error("Ошибка при изменении статуса задачи:", error);
      }
    },
  },
};
</script>

<style scoped>
.v-card {
  margin-bottom: 20px;
}
</style>
