import { useRegisterSW } from "virtual:pwa-register/react";

import { Button } from "@components/ui/button";
import { Card, CardContent } from "@components/ui/card";

export default function ReloadPrompt() {
  const {
    offlineReady: [offlineReady, setOfflineReady],
    needRefresh: [needRefresh, setNeedRefresh],
    updateServiceWorker,
  } = useRegisterSW({
    onRegistered(r) {
      if (r) {
        console.log("SW Registered: " + r);
      }
    },
    onRegisterError(error) {
      console.log("SW registration error", error);
    },
  });

  function close() {
    setOfflineReady(false);
    setNeedRefresh(false);
  }

  return (
    <div className="m-0 h-0 w-0 p-0">
      {(offlineReady || needRefresh) && (
        <Card className="fixed bottom-0 right-0 z-[9999] m-4">
          <CardContent className="p-4">
            <div className="mb-2">
              {offlineReady ? (
                <span>App ready to work offline</span>
              ) : (
                <span>Versi baru tersedia, klik tombol reload untuk memperbarui.</span>
              )}
            </div>
            <div className="flex items-center justify-start gap-2">
              {needRefresh && (
                <Button variant="outline" onClick={() => updateServiceWorker(true)}>
                  Reload
                </Button>
              )}
              <Button variant="destructive" onClick={close}>
                Close
              </Button>
            </div>
          </CardContent>
        </Card>
      )}
    </div>
  );
}
