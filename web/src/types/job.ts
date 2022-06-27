export enum SideBarType {
  SEARCH,
  INTERESTED,
  APPLIED,
  CREATE,
}

export interface Job {
  id: string;
  user_id: string;
  title: string;
  type: "full_time" | "part_time" | "internship";
  location: string;
  company: {
    company_id?: string;
    company_name: string;
  };
  salary: {
    min: number;
    max: number;
  };
  description: string;
  interested: string[];
  applied: string[];
  skills: string[];
  created: string;
}
