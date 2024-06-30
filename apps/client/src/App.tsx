import "./App.css";
import {
  TrashIcon,
  Pencil2Icon,
  CheckIcon,
  ClockIcon,
} from "@radix-ui/react-icons";
import { Input } from "./components/ui/input";
import { Label } from "./components/ui/label";
import { Button } from "./components/ui/button";
import { useEffect, useState } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./components/ui/select";
import { DialogComp } from "./components/DialogComp";
import { cn } from "./lib/utils";
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "./components/ui/tooltip";
import { appConfig } from "./config/app.config";

enum TodoStatus {
  todo = "To Do",
  progressStatus = "In Progress",
  doneStatus = "Done",
}

export interface Task {
  title: string;
  desc: string;
  status: string;
  CreatedAt: string;
  UpdatedAt: string;
  ID: number;
}

function App() {
  const [editOpen, setEditOpen] = useState<Task | null>(null);
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [tasks, setTasks] = useState<Task[]>([]);

  const fetchTasks = async (status?: TodoStatus) => {
    const res = await fetch(
      `${appConfig.apiBaseUrl}/tasks${status ? `?status=${status}` : ""}`
    );
    const jsonData = await res.json();
    setTasks(jsonData);
  };

  const createTask = async (title: string, desc: string) => {
    if (!title.length || !desc.length) return;
    await fetch(`${appConfig.apiBaseUrl}/tasks`, {
      method: "POST",
      body: JSON.stringify({
        title,
        desc,
        TaskStatus: TodoStatus.todo,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });

    fetchTasks();
  };

  const updateStatus = async (id: number, status: TodoStatus) => {
    const task = tasks.find((item) => item.ID === id);
    if (task) {
      await fetch(`${appConfig.apiBaseUrl}/tasks/${id}/status`, {
        method: "POST",
        body: JSON.stringify({
          Status: status,
        }),
        headers: {
          "Content-Type": "application/json",
        },
      });

      fetchTasks();
    }
  };

  const deleteTask = async (id: number) => {
    await fetch(`${appConfig.apiBaseUrl}/tasks/${id}`, {
      method: "DELETE",
    });

    fetchTasks();
  };

  const handleEditSave = async (title: string, desc: string) => {
    if (!title.length || !desc.length) return;
    if (editOpen) {
      await fetch(`${appConfig.apiBaseUrl}/tasks/${editOpen.ID}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title,
          desc,
        }),
      });

      setEditOpen(null);

      fetchTasks();
    }
  };

  const todoLists = () => {
    tasks.sort((a, b) => b.ID - a.ID);
    return tasks.map((task) => (
      <div
        key={task.ID}
        className="p-5 border-b grid gap-5 items-center grid-cols-1 md:grid-cols-taskLayout"
      >
        <div>
          <h4 className="font-semibold text-lg">{task.title}</h4>
          <p className="text-sm">{task.desc}</p>
        </div>
        <p
          className={cn("font-semibold text-end md:text-start", {
            "text-yellow-600": task.status === TodoStatus.progressStatus,
            "text-green-600": task.status === TodoStatus.doneStatus,
          })}
        >
          {task.status}
        </p>
        <div className="flex gap-1 items-center justify-end">
          {task.status === TodoStatus.todo && (
            <Tooltip>
              <TooltipTrigger>
                <Button
                  variant="outline"
                  onClick={() => {
                    updateStatus(task.ID, TodoStatus.progressStatus);
                  }}
                >
                  <ClockIcon />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>In Progress</p>
              </TooltipContent>
            </Tooltip>
          )}

          {task.status === TodoStatus.progressStatus && (
            <Tooltip>
              <TooltipTrigger>
                <Button
                  onClick={() => {
                    updateStatus(task.ID, TodoStatus.doneStatus);
                  }}
                >
                  <CheckIcon />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>Done</p>
              </TooltipContent>
            </Tooltip>
          )}
          {[TodoStatus.progressStatus, TodoStatus.todo].includes(
            task.status as TodoStatus
          ) && (
            <Tooltip>
              <TooltipTrigger>
                <Button
                  onClick={() => {
                    setEditOpen(task);
                  }}
                >
                  <Pencil2Icon />
                </Button>
              </TooltipTrigger>
              <TooltipContent>
                <p>Edit</p>
              </TooltipContent>
            </Tooltip>
          )}
          <Tooltip>
            <TooltipTrigger>
              <Button
                size="icon"
                variant="destructive"
                onClick={() => deleteTask(task.ID)}
              >
                <TrashIcon />
              </Button>
            </TooltipTrigger>
            <TooltipContent>
              <p>Delete</p>
            </TooltipContent>
          </Tooltip>
        </div>
      </div>
    ));
  };

  const renderTasks = () => {
    const ress = todoLists();

    if (ress.length) return ress;

    return (
      <div className="w-full h-36 flex justify-center items-center">
        <h4 className="text-xl">No Tasks Found</h4>
      </div>
    );
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  return (
    <>
      <DialogComp
        data={editOpen}
        handleEditSave={handleEditSave}
        hanldeClose={() => {
          setEditOpen(null);
        }}
      />
      <main className="w-11/12 lg:w-3/4 mx-auto">
        <h2 className="text-3xl my-5 text-center font-bold border-b pb-2">
          Task App
        </h2>
        <div className="flex items-center gap-5 mb-5">
          <div>
            <Label>Title: </Label>
            <Input
              value={title}
              className="border-primary"
              type="string"
              onChange={(e) => {
                setTitle(e.target.value);
              }}
            />
          </div>
          <div>
            <Label>Description: </Label>
            <Input
              value={desc}
              onChange={(e) => setDesc(e.target.value)}
              className="border-primary"
            ></Input>
          </div>
          <Button
            className="self-end"
            onClick={() => {
              createTask(title, desc);
              setTitle("");
              setDesc("");
            }}
            disabled={!title.length || !desc.length}
          >
            Add
          </Button>
        </div>
        <div className="border border-primary">
          <div className="bg-primary py-2 px-4 flex justify-between items-center">
            <h3 className="text-2xl text-primary-foreground">Tasks</h3>
            <Select
              defaultValue=" "
              onValueChange={(val) => fetchTasks(val as TodoStatus)}
            >
              <SelectTrigger className="w-[180px]">
                <SelectValue className="text-primary" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value=" ">All</SelectItem>
                <SelectItem value="To Do">To Do</SelectItem>
                <SelectItem value="In Progress">In Progress</SelectItem>
                <SelectItem value="Done">Done</SelectItem>
              </SelectContent>
            </Select>
          </div>
          <div>{renderTasks()}</div>
        </div>
      </main>
    </>
  );
}

export default App;
