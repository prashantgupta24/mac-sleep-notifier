#include <ctype.h>
#include <stdlib.h>
#include <stdio.h>

#include <mach/mach_port.h>
#include <mach/mach_interface.h>
#include <mach/mach_init.h>

#include <IOKit/pwr_mgt/IOPMLib.h>
#include <IOKit/IOMessage.h>

io_connect_t root_port; // a reference to the Root Power Domain IOService
// notification port allocated by IORegisterForSystemPower
IONotificationPortRef notifyPortRef;
// notifier object, used to deregister later
io_object_t notifierObject;
// this parameter is passed to the callback
void *refCon;
CFRunLoopRef runLoop;

void MySleepCallBack(void *refCon, io_service_t service, natural_t messageType, void *messageArgument)
{
    // printf("messageType %08lx, arg %08lx\n",
    //        (long unsigned int)messageType,
    //        (long unsigned int)messageArgument);
    switch (messageType)
    {
    case kIOMessageCanSystemSleep:
        if (CanSleep())
        {
            //printf("can sleep");
            IOAllowPowerChange(root_port, (long)messageArgument);
        }
        else
        {
            IOCancelPowerChange(root_port, (long)messageArgument);
        }
        break;
    case kIOMessageSystemWillSleep:
        WillSleep();
        //printf("sleeping");
        IOAllowPowerChange(root_port, (long)messageArgument);
        break;
    case kIOMessageSystemWillPowerOn:
        WillWake();
        //printf("powering on");
        break;
    case kIOMessageSystemHasPoweredOn:
        break;
    default:
        break;
    }
}

void registerNotifications()
{
    root_port = IORegisterForSystemPower(refCon, &notifyPortRef, MySleepCallBack, &notifierObject);
    if (root_port == 0)
    {
        printf("IORegisterForSystemPower failed\n");
    }

    CFRunLoopAddSource(CFRunLoopGetCurrent(),
                       IONotificationPortGetRunLoopSource(notifyPortRef), kCFRunLoopCommonModes);
    
    runLoop = CFRunLoopGetCurrent();
    CFRunLoopRun();
}

void unregisterNotifications()
{
    CFRunLoopRemoveSource(runLoop,
                          IONotificationPortGetRunLoopSource(notifyPortRef),
                          kCFRunLoopCommonModes);

    IODeregisterForSystemPower(&notifierObject);

    IOServiceClose(root_port);

    IONotificationPortDestroy(notifyPortRef);

    CFRunLoopStop(runLoop);
}