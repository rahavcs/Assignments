function sendNotification(notificationService, message) {
    const status = notificationService.send(message);
    if (status) {
      return "Notification Sent";
    } else {
      return "Failed to Send";
    }
  }
  
  module.exports = { sendNotification };
  