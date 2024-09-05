import { Card, CardHeader, CardContent } from "@/components/ui/card"
import { Header } from "./Header"
import { Sidebar } from "./Sidebar"
import { UsersIcon } from "./icons/UsersIcon"

export function Dashboard() {
  const projects = [
    {
      id: "p1",
      name: "Acme Website",
      thumbnail: "/placeholder.svg",
      createdAt: "2023-04-15",
      collaborators: 3,
    },
    {
      id: "p2",
      name: "Mobile App Design",
      thumbnail: "/placeholder.svg",
      createdAt: "2023-06-01",
      collaborators: 5,
    },
    {
      id: "p3",
      name: "Brand Identity",
      thumbnail: "/placeholder.svg",
      createdAt: "2023-02-28",
      collaborators: 2,
    },
    {
      id: "p4",
      name: "Ecommerce UI",
      thumbnail: "/placeholder.svg",
      createdAt: "2023-08-10",
      collaborators: 4,
    },
    {
      id: "p5",
      name: "Landing Page",
      thumbnail: "/placeholder.svg",
      createdAt: "2023-05-20",
      collaborators: 1,
    },
  ]

  return (
    <div className="flex min-h-screen w-full bg-muted/40">
      <Sidebar />
      <div className="flex flex-col sm:gap-4 sm:pb-4 sm:pl-14">
        <Header />
        <main className="flex-1 p-4 sm:p-6">
          <div className="grid gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
            {projects.map((project) => (
              <Card key={project.id}>
                <CardHeader>
                  <img
                    src="/placeholder.svg"
                    alt={project.name}
                    width={400}
                    height={225}
                    className="aspect-video w-full rounded-t-lg object-cover"
                  />
                </CardHeader>
                <CardContent className="p-4">
                  <div className="flex items-center justify-between">
                    <div className="text-lg font-medium">{project.name}</div>
                    <div className="flex items-center gap-1 text-muted-foreground">
                      <UsersIcon className="h-4 w-4" />
                      <span>{project.collaborators}</span>
                    </div>
                  </div>
                  <div className="text-sm text-muted-foreground">
                    Created on{" "}
                    <time dateTime={project.createdAt}>{new Date(project.createdAt).toLocaleDateString()}</time>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
        </main>
      </div>
    </div>
  )
}
