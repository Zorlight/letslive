"use client"

import { cn } from "@/utils/cn";
import { ClassValue } from "clsx";
import stream_img from "@/public/images/stream_thumbnail_example.jpg";
import { Hover3DBox } from "@/components/Hover3DBox";
import { useRouter } from "next/navigation";
import { User } from "@/types/user";
import LivestreamPreviewDetailView from "@/components/LivestreamPreviewDetailView";

const LivestreamPreviewView = ({
    className,
    viewers,
    title,
    category,
    tags,
    stream,
}: {
    className?: ClassValue;
    viewers: number;
    title: string;
    tags: string[];
    category?: string;
    stream: User;
}) => {
    const router = useRouter();

    return (
        <div className={cn("flex flex-col gap-2", className)}>
            <Hover3DBox
                viewers={viewers}
                showViewer={true}
                showStream={true}
                imageSrc={stream_img}
                className="h-[170px]"
                onClick={() => router.push(`/users/${stream.id}`)}
            />
            <LivestreamPreviewDetailView
                username={stream.username}
                title={title}
                category={category}
                tags={tags}
            />
        </div>
    );
};

export default LivestreamPreviewView;