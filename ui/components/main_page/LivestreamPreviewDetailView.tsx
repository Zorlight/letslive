import Image from "next/image";
import { Livestream } from "../../types/livestream";
import { User } from "../../types/user";
import { Clock, Eye } from "lucide-react";
import { datediffFromNow } from "../../utils/timeFormats";
const LivestreamPreviewDetailView = ({
    livestream,
    user,
}: {
    livestream: Livestream;
    user: User | null;
}) => {
    return (
        <div className="flex items-start gap-3">
            <div className="h-10 w-10 rounded-full overflow-hidden bg-muted flex-shrink-0">
                <Image
                    src={
                        user?.profilePicture || "https://github.com/shadcn.png"
                    }
                    alt={`${user?.username} avatar`}
                    className="w-full h-full object-cover"
                    width={40}
                    height={40}
                />
            </div>
            <div className="flex-1 min-w-0">
                <h3 className="font-semibold text-base truncate">
                    {livestream.title}
                </h3>
                <p className="text-sm text-muted-foreground truncate">
                    {user ? user.displayName ?? user.username : "Unknown"}
                </p>
                <div className="flex items-center gap-3 mt-1 text-xs text-muted-foreground">
                    <div className="flex items-center gap-1">
                        <Eye className="h-3 w-3" />
                        <span>{livestream.viewCount} views</span>
                    </div>
                    <div className="flex items-center gap-1">
                        <Clock className="h-3 w-3" />
                        <span>
                            Started at {datediffFromNow(livestream.startedAt)}{" "}
                            ago
                        </span>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default LivestreamPreviewDetailView;
