import { Job } from "@/src/types/job";
import { Dialog, Transition } from "@headlessui/react";
import React, { Fragment, useEffect, useState } from "react";

interface JobCreationFormProps {
  jobFormState: {
    open: boolean;
    mode: "create" | "edit";
    job: Job | null;
  };
  setJobFormState: React.Dispatch<
    React.SetStateAction<{
      open: boolean;
      mode: "create" | "edit";
      job: Job | null;
    }>
  >;
  onJobCreate: (j: Job) => void;
  onJobEdit: (j: Job) => void;
}

export default function JobCreationForm({
  jobFormState,
  setJobFormState,
  onJobCreate,
  onJobEdit,
}: JobCreationFormProps) {
  const [form, setForm] = useState<Job>({
    id: "",
    user_id: "",
    title: "",
    type: "full_time",
    location: "",
    company: {
      company_id: "",
      company_name: "",
    },
    salary: {
      min: 0,
      max: 0,
    },
    description: "",
    skills: [],
    interested: [],
    applied: [],
    created: "",
  });

  const onClose = () => {
    setJobFormState((o) => ({
      open: false,
      mode: o.mode,
      job: o.job,
    }));
  };

  useEffect(() => {
    if (jobFormState.open === true) {
      if (jobFormState.job !== null) {
        setForm({
          id: jobFormState.job.id,
          user_id: jobFormState.job.user_id,
          title: jobFormState.job.title,
          type: jobFormState.job.type,
          location: jobFormState.job.location,
          company: {
            company_id: jobFormState.job.company.company_id,
            company_name: jobFormState.job.company.company_name,
          },
          salary: {
            min: jobFormState.job.salary.min,
            max: jobFormState.job.salary.max,
          },
          description: jobFormState.job.description,
          skills: jobFormState.job.skills,
          interested: jobFormState.job.interested,
          applied: jobFormState.job.applied,
          created: jobFormState.job.created,
        });
      } else {
        setForm({
          id: "",
          user_id: "",
          title: "",
          type: "full_time",
          location: "",
          company: {
            company_id: "",
            company_name: "",
          },
          salary: {
            min: 0,
            max: 0,
          },
          description: "",
          skills: [],
          interested: [],
          applied: [],
          created: "",
        });
      }
    }
  }, [jobFormState.open]);

  return (
    <Transition appear show={jobFormState.open} as={Fragment}>
      <Dialog
        className="fixed inset-0 z-10 overflow-y-auto hello"
        onClose={onClose}
      >
        <div className="min-h-screen px-4 text-center">
          <Dialog.Overlay className="fixed inset-0 bg-black opacity-40" />

          {/* This element is to trick the browser into centering the modal contents. */}
          <span
            className="inline-block h-screen align-middle"
            aria-hidden="true"
          >
            &#8203;
          </span>
          <Transition.Child
            as={Fragment}
            enter="ease-out duration-300"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="ease-in duration-200"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <div
              className="inline-block w-full p-6 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl rounded-2xl"
              style={{ maxWidth: "900px" }}
            >
              <Dialog.Title
                as="h2"
                className="text-2xl font-medium leading-6 text-gray-900"
              >
                {jobFormState.mode === "edit"
                  ? "Edit Job Details"
                  : "Create a new Job Opening"}
              </Dialog.Title>
              <hr className="my-2" />
              <form
                onSubmit={(e) => {
                  e.preventDefault();
                  if (jobFormState.mode === "create") {
                    onJobCreate(form);
                  } else {
                    onJobEdit(form);
                  }
                }}
              >
                <div className="space-y-4">
                  <div>
                    <label className="text-gray-600">Job Title:</label>
                    <input
                      type="text"
                      value={form.title}
                      onChange={(e) =>
                        setForm((f) => ({ ...f, title: e.target.value }))
                      }
                      placeholder="ex Senior Software Enginner"
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    />
                  </div>
                  <div>
                    <label className="text-gray-600">Job Type:</label>
                    <input
                      type="text"
                      value={form.type}
                      placeholder="ex Full-Time or Part-Time or Internship"
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    />
                  </div>
                  <div>
                    <label className="text-gray-600">Company Name:</label>
                    <input
                      type="text"
                      value={form.company.company_name}
                      onChange={(e) =>
                        setForm((f) => ({
                          ...f,
                          company: {
                            ...f.company,
                            company_name: e.target.value,
                          },
                        }))
                      }
                      placeholder="ex Workflow"
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    />
                  </div>
                  <div>
                    <label className="text-gray-600">Location:</label>
                    <input
                      type="text"
                      value={form.location}
                      onChange={(e) =>
                        setForm((f) => ({ ...f, location: e.target.value }))
                      }
                      placeholder="ex Athens, Greece"
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    />
                  </div>
                  <div>
                    <label className="text-gray-600">Salary (yearly)</label>
                    <div className="flex items-center space-x-2">
                      <input
                        type="number"
                        value={form.salary.min}
                        onChange={(e) =>
                          setForm((f) => ({
                            ...f,
                            salary: {
                              ...f.salary,
                              min: parseInt(e.target.value, 10),
                            },
                          }))
                        }
                        placeholder="minimum"
                        className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                        required
                      />
                      <input
                        type="number"
                        value={form.salary.max}
                        onChange={(e) =>
                          setForm((f) => ({
                            ...f,
                            salary: {
                              ...f.salary,
                              max: parseInt(e.target.value, 10),
                            },
                          }))
                        }
                        placeholder="maximum"
                        className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                        required
                      />
                    </div>
                  </div>
                  <div>
                    <label className="text-gray-600">Job Description:</label>
                    <textarea
                      value={form.description}
                      onChange={(e) =>
                        setForm((f) => ({ ...f, description: e.target.value }))
                      }
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    ></textarea>
                  </div>
                  <div>
                    <label className="text-gray-600">
                      Skills Required - (space separated words)
                    </label>
                    <input
                      type="text"
                      value={form.skills.join(" ")}
                      onChange={(e) =>
                        setForm((f) => ({
                          ...f,
                          skills: e.target.value.split(" "),
                        }))
                      }
                      placeholder="Golang Docker Kubernetes helm git"
                      className="py-2 px-3 w-full border rounded-md border-gray-700 focus:outline-none"
                      required
                    />
                  </div>
                </div>
                <div className="flex justify-end space-x-4 mt-8">
                  <button
                    type="button"
                    className="btn border-2 border-purple-800 text-purple-800 hover:bg-purple-800 hover:text-white hover:border-0 py-2 px-3"
                    onClick={onClose}
                  >
                    Cancel
                  </button>
                  {jobFormState.mode === "edit" ? (
                    <button
                      type="submit"
                      className="btn bg-red-600 hover:bg-red-700 text-white py-2 px-3"
                    >
                      Delete Job Listing
                    </button>
                  ) : null}
                  <button
                    type="submit"
                    className="btn bg-purple-800 hover:bg-purple-900 text-white py-2 px-3"
                  >
                    {jobFormState.mode === "create"
                      ? "Create Job Listing"
                      : "Update Listing"}
                  </button>
                </div>
              </form>
            </div>
          </Transition.Child>
        </div>
      </Dialog>
    </Transition>
  );
}
